package main

import (
	"log"
	"net/http"
)

// Home Handler Function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}
func main() {
	//http.NewServeMux() function to init a new servermux=>equivalent to a router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("Starting server on :4000")
	// server creation
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
