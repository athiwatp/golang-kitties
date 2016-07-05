package main

import (
	"log"
	"net/http"

	"github.com/laibulle/kitties/app/kittiesleaf"

	"github.com/gorilla/mux"
)

func main() {
	kc := &kittiesleaf.KittiesController{}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/kitties", kc.Index).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
