package main

import (
	"log"
	"net/http"
)

func main() {
	//http.NewServeMux() function to init a new servermux=>equivalent to a router
	mux := http.NewServeMux()
	//create a file server which serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	//we have to strip static prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	// server creation
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
