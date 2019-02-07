package controllers

import (
	"io"
	"net/http"
)

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "GET INGREDIENTS")
}
func GetIngredient(w http.ResponseWriter, r *http.Request)    {}
func CreateIngredient(w http.ResponseWriter, r *http.Request) {}
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {}
func UpdateIngredient(w http.ResponseWriter, r *http.Request) {}
