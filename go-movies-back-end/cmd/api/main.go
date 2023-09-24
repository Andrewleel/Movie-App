package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string

}

func main() {
	// set applications config
	app := &application{
		Domain: "example.com",
	}
	// read from command line

	// connect to the database
	log.Println("Starting application on port: ", port)

	http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes());
	if err != nil {
		log.Fatal(err)
	}
}