package routes

import (
	"log"

	"recipe-book/controllers"
)

func BootstrapRecipeRoutes() {
	Router.HandleFunc("/recipes", controllers.GetRecipes).Methods("GET")
	Router.HandleFunc("/recipes/{d}", controllers.GetRecipe).Methods("GET")
	Router.HandleFunc("/recipes", controllers.CreateRecipe).Methods("POST")
	Router.HandleFunc("/recipes/{id}", controllers.DeleteRecipe).Methods("DELETE")
	Router.HandleFunc("/recipes/{id}", controllers.UpdateRecipe).Methods("PUT")

	log.Println("Recipe routes bootstrapped")
}
