package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/helpers"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	books, _ := helpers.ListBook(*RepBook, *RepAuthor, *RepBookAuth)
	json.NewEncoder(w).Encode(books)
}

func GetBookByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	book, _ := helpers.SearchByBookInput(params["name"], *RepBook, *RepAuthor, *RepBookAuth)
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
	book, _ := helpers.GetBookById(id, *RepBook, *RepAuthor, *RepBookAuth)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book[1])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b helpers.BookItem

	err := helpers.DecodeJSONBody(w, r, &b)
	CheckErr(err, w)

	fmt.Fprintf(w, "Person: %+v", b)
	fmt.Printf("****%s*****", b.Publisher)
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
	helpers.DeleteByBookID(id, *RepBook, *RepAuthor, *RepBookAuth)
	json.NewEncoder(w).Encode("Deleted")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
