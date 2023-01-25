package cookie

import (
	"net/http"
	"time"
)

// Cookie is a middleware for handling cookie.
type Cookie struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Cfg            *Config
}

// Config is the optional cookie config.
type Config struct {
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
}

// New creates a cookie getter and setter.
func New(w http.ResponseWriter, r *http.Request, cfg ...*Config) *Cookie {
	cfgX := DefaultCfg

	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]
	}

	return &Cookie{
		Request:        r,
		ResponseWriter: w,
		Cfg:            cfgX,
	}
}

// Set sets response cookie with the given name and value.
func (c *Cookie) Set(name string, value string, maxAge time.Duration) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(maxAge),
		Path:     c.Cfg.Path,
		Domain:   c.Cfg.Domain,
		HttpOnly: c.Cfg.HTTPOnly,
		Secure:   c.Cfg.Secure,
	}

	http.SetCookie(c.ResponseWriter, cookie)
}

// Get gets request cookie with the given name.
func (c *Cookie) Get(name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie.Value
}

// Del deletes response cookie with the given name.
func (c *Cookie) Del(name string) {
	c.Set(name, "", -1)
}
