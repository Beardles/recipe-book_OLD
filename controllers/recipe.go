package controllers

import (
	"encoding/json"
	"net/http"
	"recipe-book/models"
	"recipe-book/services"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := services.GetRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&recipes)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	recipe, err := services.GetRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&recipe)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	json.NewDecoder(r.Body).Decode(&recipe)
	if err := services.CreateRecipe(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write([]byte(`{ "success": true }`))
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if err := services.DeleteRecipe(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write([]byte(`{ "success": true }`))
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	var re models.Recipe
	json.NewDecoder(r.Body).Decode(&re)
	recipe, err := services.UpdateRecipe(&re, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&recipe)
}
