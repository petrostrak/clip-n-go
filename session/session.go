package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	mins, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		mins = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(mins) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = s.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Store = memstore.New()

	return session
}
