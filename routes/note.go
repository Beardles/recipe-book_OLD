package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapNoteRoutes() {
	Router.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	Router.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	Router.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")

	log.Println("Note routes bootstrapped")
}
