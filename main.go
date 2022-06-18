package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andfxx27/go-multiplication-table/src/domain"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	http.HandleFunc("/", domain.IndexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Application started on :%v", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
