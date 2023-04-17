package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(app.recoverPanic, app.logRequest, secureHeaders)

	r.Get("/", app.home)

	r.Route("/clip", func(r chi.Router) {
		r.Get("/{id}", app.showClip)

		r.Route("/create", func(r chi.Router) {
			r.Use(app.requireAuthentication)
			r.Get("/", app.createClipForm)
			r.Post("/", app.createClip)
		})
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/signup", app.signupUserForm)
		r.Post("/signup", app.signupUser)
		r.Get("/login", app.loginUserForm)
		r.Post("/login", app.loginUser)

		r.Route("/logout", func(r chi.Router) {
			r.Use(app.requireAuthentication)
			r.Post("/", app.logoutUser)
		})
	})

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Get("/static/*", http.StripPrefix("/static", fs).(http.HandlerFunc))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return app.Session.LoadAndSave(r)
}
