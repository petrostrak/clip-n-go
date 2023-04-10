package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
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
		app.serverError(w, err)
		return
	}

	// We then use the Execute() on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// any dynamic data that we may want to pass.
	if err = ts.Execute(w, nil); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showClip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
	}
	fmt.Fprintf(w, "Display a specific clip with ID: %d", id)
}

func (app *application) createClip(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

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

	http.Redirect(w, r, fmt.Sprintf("/clip?id=%d", id), http.StatusSeeOther)
}
