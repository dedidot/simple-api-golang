package controllers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"../models"
)

func ListBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}
	err := models.GetAllBook(db, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, books)
}

func OneBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var books models.Book
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneBookGetting(db, id, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		respondJSON(w, http.StatusOK, books)
		return
	}
}