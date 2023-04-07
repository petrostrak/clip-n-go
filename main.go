package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Clip 'n Go!" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Clip 'n Go!"))
}

func showClip(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific clip..."))
}

func createClip(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405) // It is possible to call WriteHeader() once per response
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new clip...")) // Write will automatically send 200 OK if no WriteHeader() called
}

func main() {
	// Use the http.NewServeMux() to initialize a new servemux, then
	// register the home() as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/clip", showClip)
	mux.HandleFunc("/clip/create", createClip)

	log.Println("Starting server on: 8080")
	// Use the http.ListenAndServe() to start a new web server. We pass in
	// the TCP network address to listen on and the servemux.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
