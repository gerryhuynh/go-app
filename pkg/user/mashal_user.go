package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

func MarshalUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Must be GET", http.StatusMethodNotAllowed)
		return
	}

	id := int32(1)
	email := "test@test.com"
	username := "testuser"
	firstName := "John"
	lastName := "Doe"
	age := int32(30)
	phoneNumber := "+1234567890"
	address := "123 Main St"
	city := "Anytown"
	country := "USA"
	postalCode := "12345"
	createdAt := time.Now()
	lastLoginAt := time.Now()
	isActive := true
	profilePicture := "https://example.com/profile.jpg"
	occupation := "Software Developer"
	company := "Tech Corp"

	user := &User{
		ID:             int(id),
		Email:          email,
		Username:       username,
		FirstName:      firstName,
		LastName:       lastName,
		Age:            int(age),
		PhoneNumber:    phoneNumber,
		Address:        address,
		City:           city,
		Country:        country,
		PostalCode:     postalCode,
		CreatedAt:      createdAt,
		LastLoginAt:    lastLoginAt,
		IsActive:       isActive,
		ProfilePicture: profilePicture,
		Occupation:     occupation,
		Company:        company,
	}

	person := &Person{
		Id:             id,
		Email:          email,
		Username:       username,
		FirstName:      firstName,
		LastName:       lastName,
		Age:            age,
		PhoneNumber:    phoneNumber,
		Address:        address,
		City:           city,
		Country:        country,
		PostalCode:     postalCode,
		CreatedAt:      createdAt.Unix(),
		LastLoginAt:    lastLoginAt.Unix(),
		IsActive:       isActive,
		ProfilePicture: profilePicture,
		Occupation:     occupation,
		Company:        company,
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		outputJSON(w, user)
	case "application/protobuf":
		outputProtobuf(w, person)
	default:
		outputJSON(w, user)
		outputProtobuf(w, person)
	}
}

func outputJSON(w http.ResponseWriter, user *User) {
	json, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling user: %v", err)
	}

	fmt.Println("\nJSON")
	fmt.Println("OUTPUT:", json)
	fmt.Println("SIZE:", len(json))
	fmt.Println("STRING:")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func outputProtobuf(w http.ResponseWriter, user *Person) {
	proto, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling person: %v", err)
	}

	fmt.Println("\nPROTOBUF")
	fmt.Println("OUTPUT:", proto)
	fmt.Println("SIZE:", len(proto))
	fmt.Println("STRING:")

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(http.StatusOK)
	w.Write(proto)
}
