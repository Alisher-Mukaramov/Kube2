package main

import (
	"github.com/Alisher-Mukaramov/Kube2/cmd/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	server := mux.NewRouter()
	server.HandleFunc("/health", handler.Health())
	srv := http.Server{
		Addr:    ":8000",
		Handler: server,
	}
	srv.ListenAndServe()
}
