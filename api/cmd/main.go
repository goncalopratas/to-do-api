package main

import (
	"api/infra"
)

func main() {
	infra.StartServer()

}

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type Message struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }

// // Handler for the `/` route
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	response := Message{
// 		Status:  "success",
// 		Message: "Welcome to my Go API!",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// // Handler for the `/hello` route
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	response := Message{
// 		Status:  "success",
// 		Message: "Hello, World!",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func main() {

// 	// Register routes
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/hello", helloHandler)

// 	// Start the HTTP server
// 	log.Println("Starting server on port 8080...")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
