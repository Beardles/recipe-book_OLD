package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router -> Exported mux router
var Router = mux.NewRouter()

func Bootstrap() {
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
	})
	Router.PathPrefix("/api/v1")
	BootstrapRecipeRoutes()
	BootstrapIngredientRoutes()
	BootstrapNoteRoutes()
	BootstrapStepRoutes()
	log.Println("Routes bootstrapped!")
}
