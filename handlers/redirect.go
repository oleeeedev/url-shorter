package handlers

import (
	"net/http"
	"time"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	urlMux.RLock()
	data, exists := urlStore[shortURL]
	urlMux.RUnlock()

	if !exists || time.Now().After(data.Expiry) {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	urlMux.Lock()
	data.AccessCount++
	urlStore[shortURL] = data
	urlMux.Unlock()

	http.Redirect(w, r, data.OriginalURL, http.StatusFound)
}
