package cookie

import (
	"net/http"
	"time"
)

type Cookie interface {
	// Set sets response cookie with the given name and value.
	Set(name string, value string, cfg ...*Config)
	// Get gets request cookie with the given name.
	Get(name string) string
	// Del deletes response cookie with the given name.
	Del(name string)
}

// cookie is a middleware for handling cookie.
type cookie struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Config is the optional cookie config.
type Config struct {
	MaxAge   time.Duration
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
	SameSite http.SameSite
}

// New creates a cookie getter and setter.
func New(w http.ResponseWriter, r *http.Request) Cookie {
	return &cookie{
		Request:        r,
		ResponseWriter: w,
	}
}

// Set sets response cookie with the given name and value.
func (c *cookie) Set(name string, value string, cfg ...*Config) {
	cfgX := DefaultCfg
	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]

		if cfgX.Path == "" {
			cfgX.Path = DefaultPath
		}
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(cfgX.MaxAge),
		Path:     cfgX.Path,
		Domain:   cfgX.Domain,
		HttpOnly: cfgX.HTTPOnly,
		Secure:   cfgX.Secure,
		SameSite: cfgX.SameSite,
	}

	http.SetCookie(c.ResponseWriter, cookie)
}

// Get gets request cookie with the given name.
func (c *cookie) Get(name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie.Value
}

// Del deletes response cookie with the given name.
func (c *cookie) Del(name string) {
	c.Set(name, "", &Config{MaxAge: -1})
}
