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
	http.HandleFunc("/register", handlers.RegisterUser)
	http.HandleFunc("/login", handlers.LoginUser)

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
