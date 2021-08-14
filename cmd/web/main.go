package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)


var (
	addr = flag.String("addr", ":4000", "HTTP network address")
)

type app struct {
	logger *logrus.Logger
}


func main() {
	flag.Parse()
	// Creating app instance
	a := &app{
		logger: logrus.New(),
	}
	// Creating router instance
	mux := http.NewServeMux()
	// Adding mapping urls
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)

	// Serving static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:
		*addr,
		Handler: mux,
	}

	// Force server to start listening requests
	log.Println("Starting server on :4000")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
