package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/helpers"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := helpers.ListBook(*RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	book, err := helpers.SearchByBookInput(params["name"], *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong Id")
		return
	}
	book, err := helpers.GetBookById(id, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book[0])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b helpers.BookItem

	err := helpers.DecodeJSONBody(w, r, &b)
	if err != nil {
		CheckErr(err, w)
		return
	}

	err = helpers.CreateBook(b, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong Id")
		return
	}
	err = helpers.DeleteByBookID(id, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Deleted")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var b helpers.BookItem

	err := helpers.DecodeJSONBody(w, r, &b)
	if err != nil {
		CheckErr(err, w)
		return
	}

	err = helpers.UpdateBook(b, *RepBook, *RepAuthor, *RepBookAuth)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Updated")
}

func BuyBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err1 := strconv.Atoi(params["id"])
	cnt, err2 := strconv.Atoi(params["cnt"])
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := helpers.BookBuy(id, cnt, *RepBook)
	if err != nil {
		CheckErr(err, w)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Buyed")
}
