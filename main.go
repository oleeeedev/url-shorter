package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shorter/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/stats/", handlers.GetURLStats)
	http.HandleFunc("/", handlers.RedirectURL)

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
