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
		Addr:         "0.0.0.0:8090",
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
