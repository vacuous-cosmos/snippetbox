package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"snippetbox.cosmos/internal/models"
)

// Define an application struct to hold the loggers for web application.
// For now we will only include fields for the two custom loggers
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	//Defining a  new commandline flag with the name 'addr',a default value of ":4000"
	//and some short help text explaining what the flag controls.The value of
	//the flag will be stored in addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Sql connections
	dsn := flag.String("dsn", "web:29082000@/snippetbox?parseTime=true", "MySQL data source name")
	//flag.Parse will read the command line value incase of no value default value will be provided
	flag.Parse()
	//using log.New to create a new log with three params first is where to output second is prefix third is format
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	//getting a database connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	//Initialize a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	// And add it to the application dependencies
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	//Initialize a new http.Server struct.We Set the addr and handler fields so
	//that the server uses the same network uses the same net addr and routes as before.
	//we init the ErrorLog fieldso that server uses customErrorLogger field in case of any errors
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on :%s", *addr)
	// server creation
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

//The openDB() function wraps sql.Open() and returns a sql.DB connection pool
//for a given dsn

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
