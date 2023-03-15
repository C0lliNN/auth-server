package main

import (
	"C0lliNN/auth-server/internal/config"
	"C0lliNN/auth-server/internal/container"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	appConfig := readConfig()
	s := container.CreateServer(appConfig)
	s.Start()
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
