package main

import (
	"brewery/api/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "brewery-api ", log.LstdFlags)
	prodHandler := handlers.NewProducts(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", prodHandler.GET)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.Use(prodHandler.MiddleWareProductsValidation)
	putRouter.HandleFunc("/{id:[0-9+]}", prodHandler.PUT)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.Use(prodHandler.MiddleWareProductsValidation)
	postRouter.HandleFunc("/", prodHandler.POST)

	log.Printf("Starting  the server on 8080\n")
	s := &http.Server{
		Addr:    ":8080",
		Handler: sm,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt) // Notify broadcast a message on the sigChan when an Interrupt is received
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown.", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
