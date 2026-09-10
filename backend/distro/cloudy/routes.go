package cloudy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	xutils "github.com/blue-monads/potatoverse/backend/utils"
	"github.com/gin-gonic/gin"
)

var tenantSlugRe = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

var reservedTenants = map[string]bool{
	"www":    true,
	"api":    true,
	"main":   true,
	"cloudy": true,
	"portal": true,
	"admin":  true,
	"static": true,
}

type signUpRequest struct {
	Tenant      string `json:"tenant" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PricingTier string `json:"pricing_tier"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type addUserRequest struct {
	Tenant      string `json:"tenant" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UType       string `json:"utype"`
	PricingTier string `json:"pricing_tier"`
	Verified    bool   `json:"verified"`
}

type resetPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

func (a *CloudyApp) registerBaseRouter(router *gin.Engine) {
	r := router.Group("/zz/cloudy")

	// public
	r.GET("/", a.handleIndex)
	r.GET("/health", a.handleHealth)
	r.GET("/sign-up", a.handleSignUpInfo)
	r.POST("/sign-up", a.handleSignUp)
	r.GET("/verify", a.handleVerify)
	r.POST("/login", a.handleLogin)

	// authed
	authed := r.Group("/", a.authMiddleware())
	authed.GET("/me", a.handleMe)
	authed.POST("/apps/:name/load", a.loadApp)
	authed.GET("/apps/:name/open", a.redirectToApp)

	// admin
	admin := authed.Group("/", a.adminMiddleware())
	admin.GET("/users", a.handleListUsers)
	admin.POST("/users", a.handleAddUser)
	admin.POST("/users/:id/reset-password", a.handleResetPassword)

	router.GET("/", a.handleIndex)
}

func (a *CloudyApp) handleIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "cloudy",
		"domain":  normalizeDomain(a.config.Domain),
		"message": "tenant edge for potatoverse",
	})
}

func (a *CloudyApp) handleHealth(c *gin.Context) {
	users, err := a.listUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	a.mu.RLock()
	loaded := len(a.subApps)
	a.mu.RUnlock()

	c.JSON(http.StatusOK, gin.H{
		"ok":             true,
		"users":          len(users),
		"tenants_loaded": loaded,
	})
}

func (a *CloudyApp) handleSignUpInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"fields":      []string{"tenant", "name", "email", "password", "pricing_tier"},
		"tenant_host": fmt.Sprintf("<tenant>.%s", normalizeDomain(a.config.Domain)),
	})
}

func (a *CloudyApp) handleSignUp(c *gin.Context) {
	var req signUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant := strings.ToLower(strings.TrimSpace(req.Tenant))
	if err := validateTenantSlug(tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if a.tenantExists(tenant) {
		c.JSON(http.StatusConflict, gin.H{"error": "tenant already exists"})
		return
	}

	passwordHash, err := hashSignupPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utype := UTypeNormal
	existing, err := a.listUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(existing) == 0 {
		utype = UTypeAdmin
	}

	user, err := a.insertUser(req.Name, req.Email, passwordHash, tenant, req.PricingTier, utype, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := a.sendVerificationEmail(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user created but failed to send verification email: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"message": "check your email to verify before logging in",
	})
}

func (a *CloudyApp) handleLogin(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.getUserByEmail(strings.TrimSpace(req.Email))
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if !xutils.VerifyPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if !user.IsVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "email not verified"})
		return
	}

	token, err := a.encodeClaim(&Claim{
		UserID:    user.ID,
		Email:     user.Email,
		UType:     user.UType,
		TenantKey: user.TenantKey,
		Purpose:   purposeAccess,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (a *CloudyApp) handleMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user":  getUser(c),
		"claim": getClaim(c),
	})
}

func (a *CloudyApp) handleVerify(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	claim, err := a.decodeClaim(token)
	if err != nil || claim.Purpose != purposeVerify {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired token"})
		return
	}

	user, err := a.getUserByID(claim.UserID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if !user.IsVerified {
		if err := a.markUserVerified(user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.IsVerified = true
	}

	a.mu.RLock()
	existing, alreadyLoaded := a.subApps[user.TenantKey]
	a.mu.RUnlock()
	if alreadyLoaded {
		_ = existing.WaitReady(30 * time.Second)
		host := fmt.Sprintf("%s.%s", user.TenantKey, normalizeDomain(a.config.Domain))
		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"host": host,
			"url":  fmt.Sprintf("http://%s:%d", host, a.config.Port),
		})
		return
	}

	adminPass, err := xutils.GenerateRandomString(20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sub, err := a.provisionTenant(user, adminPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	host := fmt.Sprintf("%s.%s", user.TenantKey, normalizeDomain(a.config.Domain))
	tenantURL := fmt.Sprintf("http://%s:%d", host, a.config.Port)
	_ = a.sendMail(user.Email, "Your Cloudy tenant is ready",
		fmt.Sprintf("Hi %s,\n\nTenant %s is ready.\nURL: %s\nAdmin user: %s\nAdmin password: %s\n",
			user.Fullname, user.TenantKey, tenantURL, user.Fullname, adminPass),
		fmt.Sprintf(`<p>Hi %s,</p><p>Tenant <strong>%s</strong> is ready.</p><p>URL: <a href="%s">%s</a></p><p>Admin user: %s<br/>Admin password: <code>%s</code></p>`,
			user.Fullname, user.TenantKey, tenantURL, tenantURL, user.Fullname, adminPass),
	)

	c.JSON(http.StatusOK, gin.H{
		"user":           user,
		"host":           host,
		"url":            tenantURL,
		"ready":          sub.IsReady(),
		"admin_password": adminPass,
	})
}

func (a *CloudyApp) handleListUsers(c *gin.Context) {
	users, err := a.listUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	out := make([]gin.H, 0, len(users))
	for _, u := range users {
		_, loaded := a.subApps[u.TenantKey]
		out = append(out, gin.H{
			"id":           u.ID,
			"fullname":     u.Fullname,
			"email":        u.Email,
			"tenant_key":   u.TenantKey,
			"utype":        u.UType,
			"pricing_tier": u.PricingTier,
			"is_verified":  u.IsVerified,
			"created_at":   u.CreatedAt,
			"updated_at":   u.UpdatedAt,
			"loaded":       loaded,
		})
	}
	c.JSON(http.StatusOK, gin.H{"users": out})
}

func (a *CloudyApp) handleAddUser(c *gin.Context) {
	var req addUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant := strings.ToLower(strings.TrimSpace(req.Tenant))
	if err := validateTenantSlug(tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if a.tenantExists(tenant) {
		c.JSON(http.StatusConflict, gin.H{"error": "tenant already exists"})
		return
	}

	utype := req.UType
	if utype == "" {
		utype = UTypeNormal
	}
	if !validUType(utype) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "utype must be admin or normal"})
		return
	}

	passwordHash, err := hashSignupPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := a.insertUser(req.Name, req.Email, passwordHash, tenant, req.PricingTier, utype, req.Verified)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !user.IsVerified {
		if err := a.sendVerificationEmail(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user created but failed to send verification email: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (a *CloudyApp) handleResetPassword(c *gin.Context) {
	var req resetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := parseIDParam(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := a.getUserByID(id)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	hash, err := hashSignupPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := a.updateUserPassword(user.ID, hash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "user_id": user.ID})
}

func (a *CloudyApp) redirectToApp(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	claim := getClaim(c)
	if claim.UType != UTypeAdmin && claim.TenantKey != name {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	if !a.tenantExists(name) {
		c.JSON(http.StatusNotFound, gin.H{"error": "tenant not found"})
		return
	}

	host := fmt.Sprintf("%s.%s", name, normalizeDomain(a.config.Domain))
	c.Redirect(http.StatusFound, fmt.Sprintf("http://%s:%d/zz/pages", host, a.config.Port))
}

func (a *CloudyApp) loadApp(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	if err := validateTenantSlug(name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claim := getClaim(c)
	if claim.UType != UTypeAdmin && claim.TenantKey != name {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	sub, err := a.ensureSubApp(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tenant": name,
		"port":   sub.Port,
		"ready":  sub.IsReady(),
	})
}

func (a *CloudyApp) tenantRouteMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := stripHostPort(c.Request.Host)
		tenant, isMain := a.parseTenantHost(host)
		if isMain || tenant == "" {
			c.Next()
			return
		}

		sub, err := a.ensureSubApp(tenant)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "unknown tenant"})
			c.Abort()
			return
		}

		target, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d", sub.Port))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.Rewrite = func(r *httputil.ProxyRequest) {
			r.SetURL(target)
			r.SetXForwarded()
			r.Out.Host = r.In.Host
		}
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func (a *CloudyApp) provisionTenant(user *User, adminPassword string) (*SubApp, error) {
	sub, err := NewSubApp(a.rootCtx, a.config, user.TenantKey, true, a.mailer)
	if err != nil {
		return nil, err
	}

	a.mu.Lock()
	if existing, ok := a.subApps[user.TenantKey]; ok {
		a.mu.Unlock()
		if err := existing.WaitReady(30 * time.Second); err != nil {
			return nil, err
		}
		return existing, nil
	}
	a.subApps[user.TenantKey] = sub
	a.mu.Unlock()

	if err := sub.Load(a.rootCtx, user.Fullname, adminPassword, user.Email); err != nil {
		a.mu.Lock()
		delete(a.subApps, user.TenantKey)
		a.mu.Unlock()
		return nil, err
	}
	return sub, nil
}

func (a *CloudyApp) ensureSubApp(name string) (*SubApp, error) {
	a.mu.RLock()
	if sub, ok := a.subApps[name]; ok {
		a.mu.RUnlock()
		if err := sub.WaitReady(30 * time.Second); err != nil {
			return nil, err
		}
		return sub, nil
	}
	a.mu.RUnlock()

	rec, err := a.getUserByTenant(name)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, fmt.Errorf("tenant %q not registered", name)
	}
	if !rec.IsVerified {
		return nil, fmt.Errorf("tenant %q is not verified", name)
	}

	dbPath := filepath.Join(a.config.WorkingDir, "tenants", name, "app.db")
	bootstrap := true
	if _, err := os.Stat(dbPath); err == nil {
		bootstrap = false
	}

	sub, err := NewSubApp(a.rootCtx, a.config, name, bootstrap, a.mailer)
	if err != nil {
		return nil, err
	}

	a.mu.Lock()
	if existing, ok := a.subApps[name]; ok {
		a.mu.Unlock()
		if err := existing.WaitReady(30 * time.Second); err != nil {
			return nil, err
		}
		return existing, nil
	}
	a.subApps[name] = sub
	a.mu.Unlock()

	if err := sub.Load(a.rootCtx, rec.Fullname, "", rec.Email); err != nil {
		a.mu.Lock()
		delete(a.subApps, name)
		a.mu.Unlock()
		return nil, err
	}

	return sub, nil
}

func (a *CloudyApp) parseTenantHost(host string) (tenant string, isMain bool) {
	base := normalizeDomain(a.config.Domain)
	if host == base || host == "localhost" || host == "127.0.0.1" {
		return "", true
	}

	suffix := "." + base
	if strings.HasSuffix(host, suffix) {
		sub := strings.TrimSuffix(host, suffix)
		if sub == "" || strings.Contains(sub, ".") {
			return "", true
		}
		if strings.HasPrefix(sub, "zz-") || reservedTenants[sub] {
			return "", true
		}
		return sub, false
	}

	if before, ok := strings.CutSuffix(host, ".localhost"); ok && before != "" && !strings.Contains(before, ".") {
		if reservedTenants[before] || strings.HasPrefix(before, "zz-") {
			return "", true
		}
		return before, false
	}

	return "", true
}

func stripHostPort(hostport string) string {
	host, _, err := net.SplitHostPort(hostport)
	if err != nil {
		return hostport
	}
	return host
}

func validateTenantSlug(name string) error {
	if name == "" {
		return fmt.Errorf("tenant is required")
	}
	if reservedTenants[name] {
		return fmt.Errorf("tenant name is reserved")
	}
	if strings.HasPrefix(name, "zz-") || strings.HasPrefix(name, "buddy-") {
		return fmt.Errorf("tenant name prefix is reserved")
	}
	if !tenantSlugRe.MatchString(name) {
		return fmt.Errorf("tenant must be lowercase alphanumeric with optional hyphens")
	}
	return nil
}

func parseIDParam(s string) (int64, error) {
	var id int64
	_, err := fmt.Sscan(s, &id)
	return id, err
}
