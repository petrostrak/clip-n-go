package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies
// for the web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.Int("addr", 8080, "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog,
		infoLog,
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", *addr),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on: %d.\n", *addr)

	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
