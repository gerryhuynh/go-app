package main

import (
	"go-app/pkg/user"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// 0.0.0.0:50051
// Loops on the socket on the socket
// for {
// 	conn, err := lis.Accept()
// 	if err != nil {
// 		log.Fatalf("Failed to accept connection: %v", err)
// 	}
// 	go handleConn(conn)
// }



// HTTP POST request to /user.UserService/GetUser
// Content-Type: application/protobuf
func main() {
	// http Handler
	http.Handle(user.UserService_GetUser_FullMethodName, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// we want to honour the Content-Type:
		p := &user.Person{
			Id: "1",
			Email: "test@example.com",
			Username: "johndoe",
			FirstName: "John",
			LastName: "Doe",
			Age: 30,
			PhoneNumber: "+1234567890",
			Address: "123 Main St",
		}

		switch r.Header.Get("Content-Type") {
		case "application/protobuf":
			out, err := proto.Marshal(p)
			if err != nil {
				http.Error(w, "Failed to marshal proto", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/protobuf")
			w.Write(out)
		case "application/json":
			out, err := protojson.Marshal(p)
			if err != nil {
				http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		default:
			http.Error(w, "Content-Type not supported", http.StatusUnsupportedMediaType)
			return
		}
	}))

	http.ListenAndServe(":50051", nil)
}
