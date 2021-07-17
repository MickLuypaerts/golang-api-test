package main

import (
	"brewery/api/handlers"
	"context"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "brewery"
)

// TODO /products to /product except with GET all
func main() {
	serverLogger := log.New(os.Stdout, "brewery-server ", log.LstdFlags)
	handlerLogger := log.New(os.Stdout, "brewery-api ", log.LstdFlags)
	dbLogger := log.New(os.Stdout, "brewery-db ", log.LstdFlags)
	prodHandler, err := handlers.NewProducts(handlerLogger, host, port, user, password, dbname, dbLogger)
	if err != nil {
		serverLogger.Fatal(err)
	}

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", prodHandler.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.ListSingle)

	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(options, nil)
	getRouter.Handle("/docs", sh)

	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.Use(prodHandler.MiddleWareProductsValidation)
	putRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.PUT)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.Use(prodHandler.MiddleWareProductsValidation)
	postRouter.HandleFunc("/products", prodHandler.POST)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.Delete)

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), gohandlers.AllowedOrigins([]string{"*"}), gohandlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}))
	//corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:5500"}), gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))

	s := &http.Server{
		Addr:     ":8080",
		Handler:  corsHandler(sm),
		ErrorLog: serverLogger,
	}
	serverLogger.Printf("Starting  the server on 8080\n")
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			serverLogger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt) // Notify broadcast a message on the sigChan when an Interrupt is received
	sig := <-sigChan
	serverLogger.Println("Received terminate, graceful shutdown.", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
