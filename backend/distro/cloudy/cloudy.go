package cloudy

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	_ "github.com/blue-monads/potatoverse/backend/distro"
	xutils "github.com/blue-monads/potatoverse/backend/utils"
	"github.com/blue-monads/potatoverse/backend/utils/qq"
	"github.com/gin-gonic/gin"
	"github.com/hako/branca"
	turso "turso.tech/database/tursogo"
)

//go:embed all:migrations
var migrationFS embed.FS

//go:embed all:pages
var pageFiles embed.FS

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

type Config struct {
	Port           int
	WorkingDir     string
	MasterSecret   string
	TursoAuthToken string
	TursoRemoteURL string

	// Turso platform API, used to provision one database per tenant.
	TursoAPIToken string
	TursoOrg      string
	TursoGroup    string
	TursoDBPrefix string

	Domain string
	SMTP   SMTPConfig
}

type CloudyApp struct {
	rootCtx context.Context
	tursoDB *turso.TursoSyncDb
	store   *Store
	config  *Config
	router  *gin.Engine
	mailer  *smtpSender
	branca  *branca.Branca
	turso   *tursoAPI

	mu      sync.RWMutex
	subApps map[string]*SubApp
}

func New(config *Config) (*CloudyApp, error) {
	ctx := context.Background()

	if err := os.MkdirAll(config.WorkingDir, 0o755); err != nil {
		return nil, err
	}

	bootstrap := true
	db, err := turso.NewTursoSyncDb(ctx, turso.TursoSyncDbConfig{
		Path:             filepath.Join(config.WorkingDir, "main.db"),
		RemoteUrl:        config.TursoRemoteURL,
		AuthToken:        config.TursoAuthToken,
		BootstrapIfEmpty: &bootstrap,
		//		Namespace:        "main",
	})
	if err != nil {
		return nil, err
	}

	pulled, err := db.Pull(ctx)
	if err != nil {
		return nil, err
	}
	if pulled {
		log.Println("Pulled latest changes from remote main.db")
	} else {
		log.Println("No changes to pull for main.db")
	}

	tapi := newTursoAPI(config)
	if tapi == nil {
		log.Println("warning: TURSO_API_TOKEN/TURSO_ORG not set, tenant databases will not be provisioned")
	}

	return &CloudyApp{
		rootCtx: ctx,
		tursoDB: db,
		config:  config,
		subApps: make(map[string]*SubApp),
		mailer:  newSMTPSender(config.SMTP),
		branca:  newBranca(config.MasterSecret),
		turso:   tapi,
	}, nil
}

// tenantRemote provisions (or looks up) the tenant's Turso database and returns
// the sync target for it.
func (a *CloudyApp) tenantRemote(tenant string) (tenantRemote, error) {
	if a.turso == nil {
		return tenantRemote{
			URL:       a.config.TursoRemoteURL,
			AuthToken: a.config.TursoAuthToken,
			Namespace: tenant,
		}, nil
	}
	return a.turso.ensureTenantDB(a.rootCtx, tenant)
}

func (a *CloudyApp) Build() error {
	if a.router != nil {
		return nil
	}

	if err := a.initStore(); err != nil {
		return err
	}
	if err := a.store.migrate(); err != nil {
		return err
	}

	router := gin.Default()
	router.Use(a.tenantRouteMW())
	a.registerBaseRouter(router)
	a.router = router
	return nil
}

func (a *CloudyApp) Run() error {
	qq.Println("@starting_cloudy")

	if err := a.Build(); err != nil {
		return err
	}

	go a.dbSyncer()

	addr := fmt.Sprintf(":%d", a.config.Port)
	log.Println("Cloudy listening on", addr, "domain", normalizeDomain(a.config.Domain))
	return a.router.Run(addr)
}

func (a *CloudyApp) dbSyncer() {
	for {
		ctx := context.Background()
		qq.Println("@remote_syncing")

		if _, err := a.tursoDB.Pull(ctx); err != nil {
			log.Println("Error pulling main.db:", err)
		}
		if err := a.tursoDB.Push(ctx); err != nil {
			log.Println("Error pushing main.db:", err)
		}

		a.mu.RLock()
		subs := make([]*SubApp, 0, len(a.subApps))
		for _, sub := range a.subApps {
			subs = append(subs, sub)
		}
		a.mu.RUnlock()

		for _, sub := range subs {
			if err := sub.Sync(ctx); err != nil {
				log.Printf("Error syncing tenant %s: %v", sub.Name, err)
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func (a *CloudyApp) listUsers() ([]*User, error) {
	return a.store.listUsers()
}

func (a *CloudyApp) getUserByTenant(tenantKey string) (*User, error) {
	return a.store.getUserByTenant(tenantKey)
}

func (a *CloudyApp) getUserByEmail(email string) (*User, error) {
	return a.store.getUserByEmail(email)
}

func (a *CloudyApp) getUserByID(id int64) (*User, error) {
	return a.store.getUserByID(id)
}

func (a *CloudyApp) insertUser(fullname, email, passwordHash, tenantKey, pricingTier, utype string, verified bool) (*User, error) {
	return a.store.insertUser(fullname, email, passwordHash, tenantKey, pricingTier, utype, verified)
}

func (a *CloudyApp) markUserVerified(id int64) error {
	return a.store.markUserVerified(id)
}

func (a *CloudyApp) updateUserPassword(id int64, passwordHash string) error {
	return a.store.updateUserPassword(id, passwordHash)
}

func (a *CloudyApp) tenantExists(tenantKey string) bool {
	return a.store.tenantExists(tenantKey)
}

func (a *CloudyApp) publicBaseURL() string {
	return fmt.Sprintf("http://%s:%d", normalizeDomain(a.config.Domain), a.config.Port)
}

func (a *CloudyApp) sendVerificationEmail(user *User) error {
	token, err := a.encodeClaim(&Claim{
		UserID:    user.ID,
		Email:     user.Email,
		UType:     user.UType,
		TenantKey: user.TenantKey,
		Purpose:   purposeVerify,
	})
	if err != nil {
		return err
	}

	verifyURL := fmt.Sprintf("%s/zz/cloudy/verify?token=%s", a.publicBaseURL(), token)
	subject := "Verify your Cloudy account"
	text := fmt.Sprintf("Hi %s,\n\nVerify your account for tenant %s:\n%s\n", user.Fullname, user.TenantKey, verifyURL)
	html := fmt.Sprintf(
		`<p>Hi %s,</p><p>Verify your account for tenant <strong>%s</strong>:</p><p><a href="%s">Verify email</a></p>`,
		user.Fullname, user.TenantKey, verifyURL,
	)

	return a.sendMail(user.Email, subject, text, html)
}

func normalizeDomain(domain string) string {
	return strings.TrimPrefix(domain, "*.")
}

func hashSignupPassword(password string) (string, error) {
	return xutils.HashPassword(password)
}
