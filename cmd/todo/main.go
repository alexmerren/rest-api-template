package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/internal/api"
)

func main() {
	port := 5000
	server := api.MakeRoutesAndStart(port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Router))
}
