package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/petrostrak/clip-n-go/pkg/forms"
	"github.com/petrostrak/clip-n-go/pkg/models"
)

type templateData struct {
	CurrentYear     int
	Form            *forms.Form
	Clip            *models.Clip
	Clips           []*models.Clip
	Flash           string
	IsAuthenticated bool
}

// This is a string-keyed map which acts as a lookup between the names of
// our custom template functions and the functions themselves.
var funcs = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() to get a slice of all filepaths with the extension
	// '.page.tmpl'.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name.
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles().
		ts, err := template.New(name).Funcs(funcs).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob() to add any 'layout' templates to the template set.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob() to add any 'partial' templates to the template set.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
