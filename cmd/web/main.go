package main

import (
	"flag"
	"fmt"
	"knikoda/snippetbox/pkg/postgres"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
}

var (
	driverName = "postgres"
	addr       = flag.String("addr", ":8080", "HTTP network address(default is :8080)")
)

type app struct {
	logger   *logrus.Logger
	snippets *postgres.SnippetModel
}

func main() {
	flag.Parse()
	// Creating an app instance

	db, err := openDB(os.Getenv("DB_URI"))

	if err != nil {
		fmt.Println("error while connecting to db")
	}

	if err = db.Ping(); err != nil {
		fmt.Println("error while pinging db")
	}

	if err = db.Close(); err != nil {
		fmt.Println("error while closing db connection")
	}

	a := &app{
		logger:   logrus.New(),
		snippets: &postgres.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: a.routes(),
	}

	// Force server to start listening requests
	log.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
