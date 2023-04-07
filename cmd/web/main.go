package main

import (
	"log"
	"net/http"
)

func main() {
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

	log.Println("Starting server on: 8080")
	// Use the http.ListenAndServe() to start a new web server. We pass in
	// the TCP network address to listen on and the servemux.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
