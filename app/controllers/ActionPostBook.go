package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"../models"
	"../utils"
)

func InputBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var project models.Book
	err := decoder.Decode(&project) 
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	project.Codes = utils.GenerateId()
	err = models.InsertBook(db, &project)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	var books models.Book
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneBookGetting(db, id, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&books) 
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	err = models.UpdateBook(db, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	 }
	respondJSON(w, http.StatusCreated, books)
}
 
func DeletedBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	var books models.Book
	id, _ := strconv.Atoi(vars["codes"])
	err := models.OneBookGetting(db, id, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = models.DeletedBook(db, &books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	 }
	respondJSON(w, http.StatusCreated, books)
}
 
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}