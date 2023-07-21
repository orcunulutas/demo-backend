package main

import (
	"fmt"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// app config
	var app application

	//read from cli

	// connect to the db
	app.Domain = "example.com"

	//start a web server
	portStr := fmt.Sprintf(":%d", port)

	err := http.ListenAndServe(portStr, app.routes())

	if err != nil {
		panic(err)
	}

}
