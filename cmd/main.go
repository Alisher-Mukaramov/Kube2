package main

import (
	"github.com/gorilla/mux"
	"kube2/cmd/handler"
	"net/http"
)

func main() {
	server := mux.NewRouter()
	server.HandleFunc("/health", handler.Health())
	srv := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: server,
	}
	srv.ListenAndServe()
}
