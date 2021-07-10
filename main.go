package main

// https://www.youtube.com/watch?v=DD3JlT_u0DM
// https://regex101.com/
import (
	"brewery/api/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "brewery-api ", log.LstdFlags)
	prodHandler := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", prodHandler)
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
