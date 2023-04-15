package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/petrostrak/clip-n-go/pkg/models/mysql"
	"github.com/petrostrak/clip-n-go/session"
)

// Define an application struct to hold the application-wide dependencies
// for the web application.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	clips         *mysql.ClipModel
	users         *mysql.UserModel
	templateCache map[string]*template.Template
	Session       *scs.SessionManager
}

func main() {

	addr := flag.Int("addr", 8080, "HTTP network address")
	dns := flag.String("dns", "web:pass@/clipngo?parseTime=true", "MySQL data source name")
	// secret := flag.String("secret", "mwY9HC+iHs993yzc9kZHKKMmPh+ipPFC", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := session.Session{
		CookieName:     "clipngo",
		CookieLifetime: "5",
		CookiePersist:  "true",
		CookieSecure:   "false",
		CookieDomain:   "localhost",
		SessionType:    db,
	}

	app := &application{
		errorLog,
		infoLog,
		&mysql.ClipModel{DB: db},
		&mysql.UserModel{DB: db},
		templateCache,
		session.InitSession(),
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *addr),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on: %d.\n", *addr)

	if err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"); err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
