package handlers

import (
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{l}
}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	w.Write([]byte("Hello World"))
}
