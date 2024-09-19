package user

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Age            int       `json:"age"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	PostalCode     string    `json:"postal_code"`
	CreatedAt      time.Time `json:"created_at"`
	LastLoginAt    time.Time `json:"last_login_at"`
	IsActive       bool      `json:"is_active"`
	ProfilePicture string    `json:"profile_picture"`
	Occupation     string    `json:"occupation"`
	Company        string    `json:"company"`
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
