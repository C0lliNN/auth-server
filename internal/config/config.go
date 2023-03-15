package config

type Config struct {
	Database struct {
		URI  string
		Name string
	}
	Token struct {
		Secret string
	}
	Server struct {
		Address string
	}
}
