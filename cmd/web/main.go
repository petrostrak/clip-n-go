package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/petrostrak/clip-n-go/pkg/models/mysql"
)

// Define an application struct to hold the application-wide dependencies
// for the web application.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	clips         *mysql.ClipModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.Int("addr", 8080, "HTTP network address")
	dns := flag.String("dns", "web:pass@/clipngo?parseTime=true", "MySQL data source name")
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

	app := &application{
		errorLog,
		infoLog,
		&mysql.ClipModel{DB: db},
		templateCache,
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", *addr),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on: %d.\n", *addr)

	if err = srv.ListenAndServe(); err != nil {
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
