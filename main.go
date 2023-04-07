package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Could not parse query parameter", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Display a specific clip with ID: %d", id)
}

func createClip(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
