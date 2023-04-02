package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("create a new snippet"))
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
func main() {
	//http.NewServeMux() function to init a new servermux=>equivalent to a router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	// server creation
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
