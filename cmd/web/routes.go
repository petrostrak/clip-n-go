package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standardMiddleware is being used for every request our app receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	r := chi.NewRouter()
	r.Get("/", app.home)
	r.Get("/clip/{id}", app.showClip)
	r.Get("/clip/create", app.createClipForm)
	r.Post("/clip/create", app.createClip)

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Get("/static/*", http.StripPrefix("/static", fs).(http.HandlerFunc))

	return standardMiddleware.Then(app.Session.LoadAndSave(r))
}
