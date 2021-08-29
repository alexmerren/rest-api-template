package main

import (
	"fmt"
	"log"
	"net/http"
	api "todo/pkg/api"
)

func main() {
	// This main function is basically starting the server.
	port := 5000
	server := api.NewServer(port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Router))
}
