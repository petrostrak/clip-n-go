package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    *sql.DB
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	mins, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		mins = 5
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
	session.Store = mysqlstore.New(s.SessionType)

	return session
}
