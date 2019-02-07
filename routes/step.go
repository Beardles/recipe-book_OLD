package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapStepRoutes() {
	s := Router.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/steps", controllers.CreateStep).Methods("POST")
	s.HandleFunc("/steps/{id}", controllers.DeleteStep).Methods("DELETE")
	s.HandleFunc("/steps/{id}", controllers.UpdateStep).Methods("PUT")

	log.Println("Step routes bootstrapped")
}
