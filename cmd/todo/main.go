package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/internal/api"
)

func main() {
	// This main function is basically starting the server.
	port := 5000
	server := api.NewServer(port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Router))
}
