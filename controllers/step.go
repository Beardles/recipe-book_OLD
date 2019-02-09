package controllers

import (
	"encoding/json"
	"net/http"
	"recipe-book/models"
	"recipe-book/services"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateStep(w http.ResponseWriter, r *http.Request) {
	var step models.Step
	json.NewDecoder(r.Body).Decode(&step)
	if err := services.CreateStep(&step); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write([]byte(`{ "success": true }`))
}

func DeleteStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if err := services.DeleteNote(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write([]byte(`{ "success": true }`))
}

func UpdateStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	var s models.Step
	json.NewDecoder(r.Body).Decode(&s)
	step, err := services.UpdateStep(&s, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&step)
}
