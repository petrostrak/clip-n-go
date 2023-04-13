package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	title := "二十一世紀編"
	content := `一、公元2005年，趙紫陽去世這一年，公民大眾維權事件達到
	了前所未有的高峰。全年八萬多起，每五分鐘一起。中國進入了民
	眾維權的時代- 鮑彤文集`
	expires := "7"

	id, err := app.clips.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/clip/%d", id), http.StatusSeeOther)
}

func (app *application) createClipForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new clip"))
}
