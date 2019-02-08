package main

import (
	"log"
	"net/http"
	"recipe-book/db"
	"recipe-book/routes"
)

func main() {
	db.Open()
	db.Bootstrap()
	routes.Bootstrap()

	defer db.Close()

	log.Fatal(http.ListenAndServe(":5000", routes.Router))
}
