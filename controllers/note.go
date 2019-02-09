package controllers

import (
	"encoding/json"
	"net/http"
	"recipe-book/models"
	"recipe-book/services"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	json.NewDecoder(r.Body).Decode(&note)
	if err := services.CreateNote(&note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write([]byte(`{ "success": true }`))
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
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

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	var n models.Note
	json.NewDecoder(r.Body).Decode(&n)
	note, err := services.UpdateNote(&n, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(&note)
}
