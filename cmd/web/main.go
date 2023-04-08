package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Define a new command-line flag, with the default value of 8080
	addr := flag.Int("addr", 8080, "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Use the http.NewServeMux() to initialize a new servemux, then
	// register the home() as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/clip", showClip)
	mux.HandleFunc("/clip/create", createClip)

	// Create a fileserver which serves files out of the "./ui/static" dir.
	fs := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip
	// the "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	infoLog.Printf("Starting server on: %d.\n", *addr)
	// Use the http.ListenAndServe() to start a new web server. We pass in
	// the TCP network address to listen on and the servemux.
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *addr), mux); err != nil {
		errorLog.Fatal(err)
	}
}
