package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"url-shorter/models"
	"url-shorter/utils"
)

var (
	userStore = make(map[string]models.User)
	userMux   sync.RWMutex
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hash

	userMux.Lock()
	if _, exists := userStore[user.Username]; exists {
		userMux.Unlock()
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}
	userStore[user.Username] = user
	userMux.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userMux.RLock()
	user, exists := userStore[creds.Username]
	userMux.RUnlock()
	if !exists || !utils.CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
