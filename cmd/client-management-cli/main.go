package main

import (
	"C0lliNN/auth-server/internal/auth"
	"C0lliNN/auth-server/internal/config"
	"C0lliNN/auth-server/internal/container"
	"context"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	clientType        = flag.Int("type", 1, "")
	clientSecret      = flag.String("secret", "", "")
	clientRedirectURI = flag.String("redirect-uri", "", "")
)

func main() {
	flag.Parse()

	appConfig := readConfig()
	authService := container.CreateAuth(appConfig)

	ctx := context.Background()
	req := auth.CreateClientRequest{
		Type:        auth.ClientType(*clientType),
		Secret:      *clientSecret,
		RedirectURI: *clientRedirectURI,
	}

	response, err := authService.CreateNewClient(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Client ID: ", response.ID)
}

func readConfig() config.Config {
	configPath := "./local.yml"
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var appConfig config.Config

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
