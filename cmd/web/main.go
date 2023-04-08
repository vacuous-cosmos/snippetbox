package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the loggers for web application.
// For now we will only include fields for the two custom loggers
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//Defining a  new commandline flag with the name 'addr',a default value of ":4000"
	//and some short help text explaining what the flag controls.The value of
	//the flag will be stored in addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	//flag.Parse will read the command line value incase of no value default value will be provided
	flag.Parse()
	//using log.New to create a new log with three params first is where to output second is prefix third is format
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	//http.NewServeMux() function to init a new servermux=>equivalent to a router
	mux := http.NewServeMux()
	//create a file server which serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	//we have to strip static prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	//Initialize a new http.Server struct.We Set the addr and handler fields so
	//that the server uses the same network uses the same net addr and routes as before.
	//we init the ErrorLog fieldso that server uses customErrorLogger field in case of any errors
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on :%s", *addr)
	// server creation
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
