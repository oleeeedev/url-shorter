package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetURLStats(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/stats/"):]

	urlMux.RLock()
	data, exists := urlStore[shortURL]
	urlMux.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := struct {
		OriginalURL string `json:"original_url"`
		AccessCount int    `json:"access_count"`
		Expiry      string `json:"expiry"`
	}{
		OriginalURL: data.OriginalURL,
		AccessCount: data.AccessCount,
		Expiry:      data.Expiry.Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}
