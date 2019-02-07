package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapIngredientRoutes() {
	Router.HandleFunc("/ingredients", controllers.GetIngredients).Methods("GET")
	Router.HandleFunc("/ingredients/{d}", controllers.GetIngredient).Methods("GET")
	Router.HandleFunc("/ingredients", controllers.CreateIngredient).Methods("POST")
	Router.HandleFunc("/ingredients/{id}", controllers.DeleteIngredient).Methods("DELETE")
	Router.HandleFunc("/ingredients/{id}", controllers.UpdateIngredient).Methods("PUT")

	log.Println("Ingredient routes bootstrapped")
}
