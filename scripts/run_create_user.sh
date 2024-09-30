DEFAULT_JSON='{
  "id": 1,
  "email": "john@example.com",
  "username": "johndoe",
  "first_name": "John",
  "last_name": "Doe",
  "age": 30,
  "phone_number": "+1234567890",
  "address": "123 Main St",
  "city": "Anytown",
  "country": "USA",
  "postal_code": "12345",
  "created_at": "2023-04-01T12:00:00Z",
  "last_login_at": "2023-04-01T12:00:00Z",
  "is_active": true,
  "profile_picture": "https://example.com/profile.jpg",
  "occupation": "Software Developer",
  "company": "Tech Corp"
}'

JSON_DATA=${1:-$DEFAULT_JSON}

kill $(lsof -ti:8080)
go run main.go &
sleep 1
curl -X POST http://localhost:8080/create-user \
  -H "Content-Type: application/json" \
  -d "$JSON_DATA"
