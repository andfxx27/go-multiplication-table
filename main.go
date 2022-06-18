package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andfxx27/go-static-site/src/domain"
)

func main() {
	fmt.Println("Hello world from go-static-site")

	http.HandleFunc("/", domain.IndexHandler)

	log.Println("Application started on :4000...")
	log.Fatalln(http.ListenAndServe("localhost:4000", nil))
}
