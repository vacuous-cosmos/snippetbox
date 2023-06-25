package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.cosmos/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{} //cache

	//use filepath.Glob() function to get a slice of all filepaths that
	//match the pattern "./ui/html/pages/*.tmpl".This will get us a slice
	//of all filepaths

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// Parse the base template file into a template set.
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}
		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}
		// Call ParseFiles() *on this template set* to add the
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map as normal...
		cache[name] = ts
	}
	return cache, nil
}
