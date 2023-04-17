package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	// Write the template to the buffer first, instead of straight to the w, and if
	// there is an error return.
	if err := ts.Execute(buf, app.addDefaultData(td, r)); err != nil {
		app.serverError(w, err)
	}

	// If no error occurred, write the contents of the buffer to the w.
	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.IsAuthenticated = app.isAuthenticated(r)
	td.CSRFToken = nosurf.Token(r)

	return td
}

// serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description to
// the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a wrapper around clientError which sends a 404 Not Found response.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// isAuthenticated returns true if the current request is from authenticated user, otherwise
// returns false.
func (app *application) isAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "authenticatedUserID")
}
