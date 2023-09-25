package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

func main() {
	// set applications config
	app := &application{
		Domain: "example.com",
	}
	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{
		DB: conn,
	}

	defer app.DB.Connection().Close()
	// start a web server

	log.Println("Starting application on port: ", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
