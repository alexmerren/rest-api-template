package main

import (
	"todo/internal/api"
)

func main() {
	port := 5000
	server := api.MakeRoutes(port)
	server.StartServer("", port)
}
