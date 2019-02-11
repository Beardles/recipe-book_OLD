package main

import (
	"log"
	"net/http"
	"os"
	"recipe-book/db"
	"recipe-book/routes"

	"github.com/gorilla/handlers"
)

func main() {
	db.Open()
	db.Bootstrap()
	routes.Bootstrap()

	defer db.Close()

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS()(handlers.LoggingHandler(os.Stdout, routes.Router))))
}
