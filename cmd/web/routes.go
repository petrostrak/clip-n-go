package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/clip", app.showClip)
	mux.HandleFunc("/clip/create", app.createClip)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware.
	return secureHeaders(mux)
}
