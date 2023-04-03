package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home Handler Function
func home(w http.ResponseWriter, r *http.Request) {
	//check if the currentl request URL path matf=ches "/"
	// otherwise send a 404 response
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl", "./ui/html/partials/nav.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// snippetView Handler Function
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display s specific snippet with ID %d", id)
}

// snippetcreate Handler Function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //check request type
		//if not post then writeheader as 405 and send a method not allowed response
		w.Header().Set("Allow", "POST")
		//calls writeHeader and write function behind the scenes
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new request"))
}
