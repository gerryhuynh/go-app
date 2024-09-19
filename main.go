package main

import (
	"encoding/json"
	"fmt"
	"go-app/pkg/user"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Protobuf
	person := &user.Person{
		Id:             "1",
		Email:          "test@example.com",
		Username:       "johndoe",
		FirstName:      "John",
		LastName:       "Doe",
		Age:            30,
		PhoneNumber:    "+1234567890",
		Address:        "123 Main St",
		City:           "New York",
		Country:        "USA",
		PostalCode:     "10001",
		CreatedAt:      1648656000, // Unix timestamp for a sample date
		LastLoginAt:    1648742400, // Unix timestamp for a sample date
		IsActive:       true,
		ProfilePicture: "https://example.com/profile.jpg",
		Occupation:     "Software Engineer",
		Company:        "Tech Corp",
	}

	// JSON
	user := &user.User{
		ID:             "1",
		Email:          "test@example.com",
		Username:       "johndoe",
		FirstName:      "John",
		LastName:       "Doe",
		Age:            30,
		PhoneNumber:    "+1234567890",
		Address:        "123 Main St",
		City:           "New York",
		Country:        "USA",
		PostalCode:     "10001",
		CreatedAt:      time.Now().Add(-24 * time.Hour), // 1 day ago
		LastLoginAt:    time.Now(),
		IsActive:       true,
		ProfilePicture: "https://example.com/profile.jpg",
		Occupation:     "Software Engineer",
		Company:        "Tech Corp",
	}

	jsonOutput, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println("JSON Output:", string(jsonOutput))
	fmt.Println("JSON Size:", len(jsonOutput))


	protoOutput, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to Protobuf:", err)
		return
	}

	fmt.Println("Protobuf Output:", string(protoOutput))
	fmt.Println("Protobuf Size:", len(protoOutput))
}
