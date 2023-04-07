package main

import (
	"fmt"
	"html/template"
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

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/main.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() to read the template file into a
	// template set.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not parse template file", http.StatusInternalServerError)
		return
	}

	// We then use the Execute() on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// any dynamic data that we may want to pass.
	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not execute template set", http.StatusInternalServerError)
	}
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
