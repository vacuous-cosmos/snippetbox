package main

import (
	"fmt"
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
	w.Write([]byte("Hello from snippetbox"))
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
