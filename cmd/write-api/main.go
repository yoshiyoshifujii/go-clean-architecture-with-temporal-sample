package main

import (
	"log"
	"yoshiyoshifujii/go-clean-architecture-with-temporal-sample/internal/interface/api"
)

func main() {
	server := api.NewServer()
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
