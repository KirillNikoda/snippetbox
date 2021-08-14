package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	addr = flag.String("addr", ":4000", "HTTP network address(default is :4000)")
)

type app struct {
	logger *logrus.Logger
}

func main() {
	flag.Parse()
	// Creating an app instance
	a := &app{
		logger: logrus.New(),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: a.routes(),
	}

	// Force server to start listening requests
	log.Println("Starting server on :4000")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
