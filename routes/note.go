package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapNoteRoutes() {
	s := Router.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	s.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	s.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")

	log.Println("Note routes bootstrapped")
}
