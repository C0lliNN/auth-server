package main

import (
	"C0lliNN/auth-server/internal/auth"
	"C0lliNN/auth-server/internal/container"
	"context"
	"flag"
	"fmt"
)

var (
	clientType        = flag.Int("type", 1, "")
	clientSecret      = flag.String("secret", "", "")
	clientRedirectURI = flag.String("redirect-uri", "", "")
)

func main() {
	flag.Parse()

	a := container.CreateAuth()

	ctx := context.Background()
	req := auth.CreateClientRequest{
		Type:        auth.ClientType(*clientType),
		Secret:      *clientSecret,
		RedirectURI: *clientRedirectURI,
	}

	response, err := a.CreateNewClient(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Client ID: ", response.ID)
}
