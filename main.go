package main

import (
	"log"
	"net/http"
	"recipe-book/db"
	"recipe-book/routes"
)

func main() {
	db.Bootstrap()
	routes.Bootstrap()

	log.Fatal(http.ListenAndServe(":5000", routes.Router))
}
