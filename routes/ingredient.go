package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapIngredientRoutes() {
	s := Router.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/ingredients", controllers.GetIngredients).Methods("GET")
	s.HandleFunc("/ingredients/{id}", controllers.GetIngredient).Methods("GET")
	s.HandleFunc("/ingredients", controllers.CreateIngredient).Methods("POST")
	s.HandleFunc("/ingredients/{id}", controllers.DeleteIngredient).Methods("DELETE")
	s.HandleFunc("/ingredients/{id}", controllers.UpdateIngredient).Methods("PUT")

	log.Println("Ingredient routes bootstrapped")
}
