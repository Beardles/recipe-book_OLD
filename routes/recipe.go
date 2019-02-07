package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapRecipeRoutes() {
	s := Router.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/recipes", controllers.GetRecipes).Methods("GET")
	s.HandleFunc("/recipes/{d}", controllers.GetRecipe).Methods("GET")
	s.HandleFunc("/recipes", controllers.CreateRecipe).Methods("POST")
	s.HandleFunc("/recipes/{id}", controllers.DeleteRecipe).Methods("DELETE")
	s.HandleFunc("/recipes/{id}", controllers.UpdateRecipe).Methods("PUT")

	log.Println("Recipe routes bootstrapped")
}
