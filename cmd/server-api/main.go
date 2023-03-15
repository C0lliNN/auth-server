package main

import (
	"C0lliNN/auth-server/internal/container"
	"C0lliNN/auth-server/internal/server"
)

func main() {
	s := server.NewServer(container.CreateAuth())
	s.Start()
}
