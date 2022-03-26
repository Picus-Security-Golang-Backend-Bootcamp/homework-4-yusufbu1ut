package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/helpers"
	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors, err := helpers.ListAuth(*RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	author, err := helpers.SearchByAuthorInput(params["name"], *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong Using")
		return
	}
	author, err := helpers.GetAuthById(id, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author[0])
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var a helpers.AuthorItem

	err := helpers.DecodeJSONBody(w, r, &a)
	if err != nil {
		CheckErr(err, w)
		return
	}
	err = helpers.CreateAuthor(a, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong Id")
		return
	}
	err = helpers.DeleteByAuthorID(id, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Deleted")
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var a helpers.AuthorItem
	err := helpers.DecodeJSONBody(w, r, &a)
	if err != nil {
		CheckErr(err, w)
		return
	}
	err = helpers.UpdateAuthor(a, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Updated")
}
