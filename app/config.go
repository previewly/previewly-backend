package app

import (
	"flag"
)

const (
	defaultHost     = "localhost"
	defaultHTTPHost = "http://" + defaultHost
)

type Postgres struct {
	Port     int
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
	ImageHost string
}

func newConfig() Config {
	var (
		imageHostFlag string

		dbHostFlag         string
		dbPortFlag         int
		dbNameFlag         string
		dbUserNameFlag     string
		dbUserPasswordFlag string

		apiURLFlag string
	)

	flag.StringVar(&imageHostFlag, "image-host", defaultHTTPHost+":8000", "Image host")

	flag.StringVar(&dbHostFlag, "db-host", defaultHost, "Database host")
	flag.IntVar(&dbPortFlag, "db-port", 5432, "Database port")
	flag.StringVar(&dbNameFlag, "db-name", "wsw", "Database name")
	flag.StringVar(&dbUserNameFlag, "db-user-name", "wsw", "Database user name")
	flag.StringVar(&dbUserPasswordFlag, "db-user-password", "wsw", "Database user password")

	flag.StringVar(&apiURLFlag, "api-url", defaultHTTPHost+":7171/api", "Api url")

	flag.Parse()

	config := Config{
		App: AppConfig{
			ImageHost: imageHostFlag,
		},
		Postgres: Postgres{
			Host:     dbHostFlag,
			Port:     dbPortFlag,
			DB:       dbNameFlag,
			User:     dbUserNameFlag,
			Password: dbUserPasswordFlag,
		},
		API: API{
			BaseURL: apiURLFlag,
		},
	}
	// utils.D(config)
	return config
}
