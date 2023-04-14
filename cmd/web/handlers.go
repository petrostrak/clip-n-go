package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/petrostrak/clip-n-go/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	clip, err := app.clips.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Clips: clip})
}

func (app *application) showClip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.notFound(w)
	}

	clip, err := app.clips.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Clip: clip})
}

func (app *application) createClip(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	// If there are any validation errors, re-display the create.page.tmpl
	// passing in the validation errors and previously submitted r.PostForm data.
	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.clips.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/clip/%d", id), http.StatusSeeOther)
}

func (app *application) createClipForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}
