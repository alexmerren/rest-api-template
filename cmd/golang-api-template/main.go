package main

import (
	"fmt"
	"golang-api-template/internal/server"
)

const (
	hostname = "127.0.0.1"
	port     = 5000
)

func main() {
	s := server.NewServer()
	fmt.Printf("starting server on %s:%d\n", hostname, port)
	s.StartServer(hostname, port)
}
