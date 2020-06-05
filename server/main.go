package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func receive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/dev/receive").Subrouter()
	api.HandleFunc("", receive).Methods(http.MethodGet)
	api.HandleFunc("", receive).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
