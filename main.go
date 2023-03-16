package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"

	"C0lliNN/auth-server/op"
	"C0lliNN/auth-server/storage"
)

type Config struct {
	OpenConnectID struct {
		Key             string
		Issuer          string
		WebClientSecret string
		APIClientSecret string
	}
	HTTP struct {
		Address string
	}
}

func readConfig() Config {
	configPath := "./local.yml"
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var appConfig Config

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&appConfig); err != nil {
		panic(err)
	}

	return appConfig
}

func main() {
	config := readConfig()

	storage.RegisterClients(
		storage.NativeClient("native"),
		storage.WebClient("web", config.OpenConnectID.WebClientSecret),
		storage.WebClient("api", config.OpenConnectID.APIClientSecret),
	)

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore(config.OpenConnectID.Issuer))

	router := op.NewServer(config.OpenConnectID.Issuer, storage, config.OpenConnectID.Key)

	server := &http.Server{
		Addr:    config.HTTP.Address,
		Handler: router,
	}
	log.Printf("server listening on http://%s", config.HTTP.Address)
	log.Println("press ctrl+c to stop")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
