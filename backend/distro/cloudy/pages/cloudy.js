window.Cloudy = {
  tokenKey: "cloudy_token",
  pages: "/zz/cloudy/pages",

  token() {
    return localStorage.getItem(this.tokenKey);
  },

  setToken(t) {
    localStorage.setItem(this.tokenKey, t);
  },

  page(name) {
    return this.pages + "/" + name;
  },

  signOut() {
    localStorage.removeItem(this.tokenKey);
    location.href = this.page("login.html");
  },

  headers(json) {
    const h = { Authorization: "Bearer " + this.token() };
    if (json) h["Content-Type"] = "application/json";
    return h;
  },

  async api(path, opts) {
    const res = await fetch(path, opts);
    const data = await res.json().catch(() => ({}));
    if (!res.ok) throw new Error(data.error || res.statusText || "request failed");
    return data;
  },

  afterLogin(user, token) {
    this.setToken(token);
    location.href = this.page(user && user.utype === "admin" ? "admin-portal.html" : "portal.html");
  },

  tenantURL(session, tenant) {
    const host = tenant + "." + session.domain;
    return "http://" + host + ":" + session.port;
  },

  async requireSession(opts) {
    opts = opts || {};
    if (!this.token()) {
      location.href = this.page("login.html");
      return null;
    }
    try {
      const data = await this.api("/zz/cloudy/me", { headers: this.headers() });
      const utype = data.user && data.user.utype;
      if (opts.admin && utype !== "admin") {
        location.href = this.page("portal.html");
        return null;
      }
      return data;
    } catch (err) {
      this.signOut();
      return null;
    }
  },
};
