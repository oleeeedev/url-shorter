package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
	"url-shorter/models"
	"url-shorter/utils"
)

var (
	urlStore = make(map[string]models.URLData)
	mux      sync.RWMutex
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL       string `json:"url"`
		ExpiryIn  int    `json:"expiry_in"`  // Expiry time in minutes
		CustomURL string `json:"custom_url"` // Optional custom short URL
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL := request.CustomURL
	if shortURL == "" {
		shortURL = utils.GenerateShortURL()
	}

	expiry := time.Now().Add(time.Duration(request.ExpiryIn) * time.Minute)
	mux.Lock()
	if _, exists := urlStore[shortURL]; exists {
		mux.Unlock()
		http.Error(w, "Custom URL already exists", http.StatusBadRequest)
		return
	}
	urlStore[shortURL] = models.URLData{OriginalURL: request.URL, Expiry: expiry}
	mux.Unlock()

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortURL),
	}
	json.NewEncoder(w).Encode(response)
}
