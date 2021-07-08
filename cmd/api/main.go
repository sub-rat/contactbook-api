package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sub-rat/contactbook-api/internal/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}
	srv := server.GetServer()
	port := os.Getenv("PORT")
	fmt.Printf("Starting Server at %s", port)
	srv.Run()
}
