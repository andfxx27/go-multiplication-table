package main

import (
	"log"
	"net/http"

	"github.com/andfxx27/go-multiplication-table/src/domain"
)

func main() {
	http.HandleFunc("/", domain.IndexHandler)

	log.Println("Application started on :4000...")
	log.Fatalln(http.ListenAndServe("localhost:4000", nil))
}
