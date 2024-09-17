package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
}

type Response struct {
	User    User   `json:"user"`
	Message string `json:"message"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Must be POST", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := Response{
		User:    newUser,
		Message: "User created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
