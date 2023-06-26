package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.cosmos/internal/models"
)

// Dependency injection is done by defining b= afunction against a particular *application
// Home Handler Function

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if the currentl request URL path matf=ches "/"
	// otherwise send a 404 response
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	//updating the newtemplate Data to get a init struct with current year
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.tmpl", data)
}

// snippetView Handler Function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	//Use the SnippetModels objects Get methos to retrieve the data
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.tmpl", data)
}

// snippetcreate Handler Function
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //check request type
		//if not post then writeheader as 405 and send a method not allowed response
		w.Header().Set("Allow", "POST")
		//calls writeHeader and write function behind the scenes
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	//using dummy variables for now
	title := "o snail"
	content := "O snail\n Climb Mount Fuki,\n But slowly,slowly!\n\n- Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
