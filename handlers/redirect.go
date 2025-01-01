package handlers

import (
	"net/http"
	"time"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	mux.RLock()
	data, exists := urlStore[shortURL]
	mux.RUnlock()

	if !exists || time.Now().After(data.Expiry) {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	mux.Lock()
	data.AccessCount++
	urlStore[shortURL] = data
	mux.Unlock()

	http.Redirect(w, r, data.OriginalURL, http.StatusFound)
}
