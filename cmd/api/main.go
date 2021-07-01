package main

import (
	"fmt"

	"github.com/sub-rat/contactbook-api/internal/server"
)

func main() {
	srv := server.GetServer()
	fmt.Println("Starting Server....")
	srv.Run()
}
