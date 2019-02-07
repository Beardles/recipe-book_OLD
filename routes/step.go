package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapStepRoutes() {
	Router.HandleFunc("/steps", controllers.CreateStep).Methods("POST")
	Router.HandleFunc("/steps/{id}", controllers.DeleteStep).Methods("DELETE")
	Router.HandleFunc("/steps/{id}", controllers.UpdateStep).Methods("PUT")

	log.Println("Step routes bootstrapped")
}
