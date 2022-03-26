package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/helpers"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/repositories"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	RepBook     *repositories.BookRepository
	RepAuthor   *repositories.AuthorRepository
	RepBookAuth *repositories.BookAuthRepository
)

// IndexRouting makes routings for handlers
func IndexRouting(bk *repositories.BookRepository, aut *repositories.AuthorRepository, bkAut *repositories.BookAuthRepository) *mux.Router {
	RepBook, RepAuthor, RepBookAuth = bk, aut, bkAut
	r := mux.NewRouter()
	CORSOptions()
	r.Use(loggingMiddleware)

	//? <--------------------Book-------------------------->
	book := r.PathPrefix("/books").Subrouter()

	book.HandleFunc("/", GetAllBooks).Methods("GET")
	book.HandleFunc("", GetAllBooks).Methods("GET")

	book.HandleFunc("/search={name}", GetBookByName).Methods("GET")

	book.HandleFunc("/id={id:[0-9]+}", GetBookById).Methods("GET")
	book.HandleFunc("/{id:[0-9]+}", GetBookById).Methods("GET")

	book.HandleFunc("/id={id:[0-9]+}", DeleteBook).Methods("DELETE")
	book.HandleFunc("/{id:[0-9]+}", DeleteBook).Methods("DELETE")

	book.HandleFunc("/create", CreateBook).Methods("POST")

	book.HandleFunc("/update", UpdateBook).Methods("PATCH")

	book.HandleFunc("/{id:[0-9]+}/count={cnt:[0-9]+}", BuyBook).Methods("POST")

	//? <--------------------Author-------------------------->
	author := r.PathPrefix("/authors").Subrouter()

	author.HandleFunc("/", GetAllAuthors).Methods("GET")
	author.HandleFunc("", GetAllAuthors).Methods("GET")

	author.HandleFunc("/search={name}", GetAuthorByName).Methods("GET")

	author.HandleFunc("/{id:[0-9]+}", GetAuthorById).Methods("GET")
	author.HandleFunc("/id={id:[0-9]+}", GetAuthorById).Methods("GET")

	author.HandleFunc("/{id:[0-9]+}", DeleteAuthor).Methods("DELETE")
	author.HandleFunc("/id={id:[0-9]+}", DeleteAuthor).Methods("DELETE")

	author.HandleFunc("/create", CreateAuthor).Methods("POST")

	author.HandleFunc("/update", UpdateAuthor).Methods("PATCH")

	return r
}

func CORSOptions() {
	handlers.AllowedOrigins([]string{"localhost:8090"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH", "DELETE"})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Query())
		next.ServeHTTP(w, r)
	})
}

// CheckErr checks MalformedRequest errors and makes responses
func CheckErr(err error, w http.ResponseWriter) {
	var mr *helpers.MalformedRequest
	if errors.As(err, &mr) {
		http.Error(w, mr.Msg, mr.Status)
	} else {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// authenticationMiddleware didnt use yet
// func authenticationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if strings.HasPrefix(r.URL.Path, "/products") {
// 			if token != "" {
// 				next.ServeHTTP(w, r)
// 			} else {
// 				http.Error(w, "Token not found", http.StatusUnauthorized)
// 			}
// 		} else {
// 			next.ServeHTTP(w, r)
// 		}

// 	})
// }
