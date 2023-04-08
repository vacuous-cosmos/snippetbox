package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//Defining a  new commandline flag with the name 'addr',a default value of ":4000"
	//and some short help text explaining what the flag controls.The value of
	//the flag will be stored in addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	//flag.Parse will read the command line value incase of no value default value will be provided
	flag.Parse()
	//http.NewServeMux() function to init a new servermux=>equivalent to a router
	mux := http.NewServeMux()
	//create a file server which serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	//we have to strip static prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Printf("Starting server on :%s", *addr)
	// server creation
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
