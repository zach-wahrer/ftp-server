package main

import (
	"internal/server"
)

func main() {
	server.Listen("localhost", "8000")
}
