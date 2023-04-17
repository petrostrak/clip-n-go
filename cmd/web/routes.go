package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// standardMiddleware is being used for every request our app receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	authMiddleware := alice.New(app.requireAuthentication)

	r := chi.NewRouter()
	r.Get("/", app.home)
	r.Get("/clip/{id}", app.showClip)
	r.Use(app.requireAuthentication)
	r.Get("/clip/create", authMiddleware.ThenFunc(app.createClipForm).(http.HandlerFunc))
	r.Post("/clip/create", authMiddleware.ThenFunc(app.createClip).(http.HandlerFunc))

	r.Get("/user/signup", app.signupUserForm)
	r.Post("/user/signup", app.signupUser)
	r.Get("/user/login", app.loginUserForm)
	r.Post("/user/login", app.loginUser)
	r.Post("/user/logout", authMiddleware.ThenFunc(app.logoutUser).(http.HandlerFunc))

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Get("/static/*", http.StripPrefix("/static", fs).(http.HandlerFunc))

	return standardMiddleware.Then(app.Session.LoadAndSave(r))
}
