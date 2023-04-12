package main

import (
	"html/template"
	"path/filepath"

	"github.com/petrostrak/clip-n-go/pkg/models"
)

type templateData struct {
	CurrentYear int
	Clip        *models.Clip
	Clips       []*models.Clip
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

		// Parse page template
		ts, err := template.ParseFiles(page)
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
