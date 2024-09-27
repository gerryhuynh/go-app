DEFAULT_JSON='{"id": "1", "email": "john@example.com"}'
JSON_DATA=${1:-$DEFAULT_JSON}

kill $(lsof -ti:8080)
go run main.go &
sleep 1
curl -X POST http://localhost:8080/create-user \
  -H "Content-Type: application/json" \
  -d "$JSON_DATA"
