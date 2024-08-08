package app

import (
	"fmt"
	"os"
)

const (
	defaultHost     = "localhost"
	defaultHTTPHost = "http://" + defaultHost
)

type Postgres struct {
	Port     string
	Host     string
	DB       string
	User     string
	Password string
}

type API struct {
	BaseURL string
}

type Config struct {
	App      AppConfig
	Postgres Postgres
	API      API
}

type AppConfig struct {
	Host Hosts
}

type Hosts struct {
	Images string
}

func newConfig() Config {
	fmt.Println()
	fmt.Println("You could use this environment variables: DBPort, DBHost, DBName, DBUser, DBPassword, ApiUrl, ImageHost")
	fmt.Println()

	config := Config{
		App: AppConfig{
			Host: Hosts{
				Images: getEnvVariable("ImageHost", defaultHTTPHost+":8000"),
			},
		},
		Postgres: Postgres{
			Port:     getEnvVariable("DBPort", "5432"),
			Host:     getEnvVariable("DBHost", defaultHost),
			DB:       getEnvVariable("DBName", "wsw"),
			User:     getEnvVariable("DBUser", "wsw"),
			Password: getEnvVariable("DBPassword", "wsw"),
		},
		API: API{
			BaseURL: getEnvVariable("ApiUrl", defaultHTTPHost+":7171/api"),
		},
	}
	// utils.D(config)
	return config
}

func getEnvVariable(variableName string, defaultValue string) string {
	if os.Getenv(variableName) != "" {
		return os.Getenv(variableName)
	}

	return defaultValue
}
