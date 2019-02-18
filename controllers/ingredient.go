package controllers

import (
	"encoding/json"
	"net/http"
	"recipe-book/models"
	"recipe-book/services"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := services.GetIngredients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredientDTOList := services.BuildIngredientDTOList(ingredients)

	json.NewEncoder(w).Encode(&ingredientDTOList)
}

func GetIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredient, err := services.GetIngredient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredientDTO := services.BuildIngredientDTO(ingredient)

	json.NewEncoder(w).Encode(&ingredientDTO)

}

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient models.Ingredient
	json.NewDecoder(r.Body).Decode(&ingredient)

	ingredients, err := services.CreateIngredient(&ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredientDTOList := services.BuildIngredientDTOList(ingredients)

	// w.Write([]byte(`{ "success": true }`))
	json.NewEncoder(w).Encode(&ingredientDTOList)
}

func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredients, err := services.DeleteIngredient(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	ingredientDTOList := services.BuildIngredientDTOList(ingredients)

	// w.Write([]byte(`{ "success": true }`))
	json.NewEncoder(w).Encode(&ingredientDTOList)
}

func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	var i models.Ingredient
	json.NewDecoder(r.Body).Decode(&i)
	ingredient, err := services.UpdateIngredient(&i, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&ingredient)
}
