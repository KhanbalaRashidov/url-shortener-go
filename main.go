package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener-go/internal/data"
	_ "url-shortener-go/internal/data"
)

func main() {
	log.Print("Hello world sample started.")
	r := mux.NewRouter()
	redirectPath := "http://localhost:8080/r"

	fs, err := data.NewFileStore("testing.json")
	if err != nil {
		panic("unable to create filestore appropriately")
	}

	r.Handle("/", &HandleViaStruct{}).Methods("GET")
	r.Handle("/add", &AddPath{domain: redirectPath, store: &fs}).
		Methods("POST")
	r.Handle("/r/{hash}", &DeletePath{store: &fs}).Methods("DELETE")
	r.Handle("/r/{hash}", &RedirectPath{store: &fs}).Methods("GET")
	http.ListenAndServe(":8080", r)

}
