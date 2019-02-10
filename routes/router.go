package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router -> Exported mux router
var Router = mux.NewRouter()

// Bootstrap -> Func to bootrasp all routes
func Bootstrap() {
	BootstrapRecipeRoutes()
	BootstrapIngredientRoutes()
	BootstrapNoteRoutes()
	BootstrapStepRoutes()
	Router.PathPrefix("/static").Handler(http.FileServer(http.Dir("./client/build")))
	Router.Handle("/", http.FileServer(http.Dir("./client/build")))
	log.Println("Routes bootstrapped!")
}
