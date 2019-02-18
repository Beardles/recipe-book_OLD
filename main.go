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

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headersOk, originsOk, methodsOk)(handlers.LoggingHandler(os.Stdout, routes.Router))))
}
