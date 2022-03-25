package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/handlers"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-yusufbu1ut/internal/helpers"
)

func main() {
	log.Println("Starting Server...")
	bookRepository, authorRepository, bookAuthorRepository := helpers.ConnectDB()
	router := handlers.IndexRouting(bookRepository, authorRepository, bookAuthorRepository)
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)

}

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting Down...")
	os.Exit(0)
}

// var (
// 	repositories struct {
// 		repBook     *book.BookRepository
// 		repbookAuth *bookAuthor.BookAuthRepository
// 		repAuth     *author.AuthorRepository
// 		router      mux.Router
// 	}
// )

// func init() {
// 	repositories.repBook, repositories.repAuth, repositories.repbookAuth = csvToDB.ToConnectDB()
// }

// func main() {
// 	router := mux.NewRouter()
//
// 	router.Use(loggingMiddleware)
// 	router.Use(authenticationMiddleware)

// 	s := router.PathPrefix("/books").Subrouter()
// 	//s.HandleFunc("name",HandlerBookName)
// 	s.HandleFunc("/id/{id:[0-9]+}", HandlerBookId)

// 	// Run our server in a goroutine so that it doesn't block.
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil {
// 			log.Println(err)
// 		}
// 	}()

// 	ShutdownServer(srv, time.Second*10)
// }

// func HandlerBookId(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-type", "application/json")
// 	d := ApiResponse{
// 		Data: vars["id"],
// 	}
// 	resp, _ := json.Marshal(d)
// 	w.Write(resp)
// }

// func authenticationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		token := r.Header.Get("Authorization")
// 		if strings.HasPrefix(r.URL.Path, "/books") {
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

// type ApiResponse struct {
// 	Data interface{} `json:"data"`
// }
