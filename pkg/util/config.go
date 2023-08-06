package util

type Config struct {
	SecretKey string
	Port      string
}

func LoadConfig() Config {
	return Config{
		SecretKey: "secret",
		Port:      "8080",
	}
}
