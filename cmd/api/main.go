package main

import (
	"fmt"

	"github.com/sub-rat/contactbook-api/internal/server"
)

// // Homepage
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message":"Welome to api"}`))
// }

// // Fetching Contact
// func getContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"name":"ram", "phone":"1232424"}`))
// }

// // Creating new contact
// func createContact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message":"successfully Created"}`))
// }

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", home).Methods(http.MethodGet)
	// r.HandleFunc("/contact", getContact).Methods(http.MethodGet)
	// r.HandleFunc("/contact", createContact).Methods(http.MethodPost)
	// log.Println("Server stated at 8080")
	// log.Fatal(http.ListenAndServe(":8081", r))
	srv := server.GetServer()
	fmt.Println("Starting Server....")
	srv.Run()
}
