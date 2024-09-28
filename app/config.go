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

type Gowitness struct {
	ScreenshotPath string
}

type Config struct {
	App       AppConfig
	Postgres  Postgres
	Gowitness Gowitness
}

type ListenHost struct {
	Host string
	Port int
}

type AppConfig struct {
	Listen ListenHost
}

func newConfig() Config {
	var (
		listenHostFlag string
		listenPortFlag int

		dbHostFlag         string
		dbPortFlag         int
		dbNameFlag         string
		dbUserNameFlag     string
		dbUserPasswordFlag string

		screenShotPath string
	)

	flag.StringVar(&listenHostFlag, "listen-host", "", "Listen host")
	flag.IntVar(&listenPortFlag, "listen-port", 8000, "Listen port")

	flag.StringVar(&dbHostFlag, "db-host", defaultHost, "Database host")
	flag.IntVar(&dbPortFlag, "db-port", 5432, "Database port")
	flag.StringVar(&dbNameFlag, "db-name", "wsw", "Database name")
	flag.StringVar(&dbUserNameFlag, "db-user-name", "wsw", "Database user name")
	flag.StringVar(&dbUserPasswordFlag, "db-user-password", "wsw", "Database user password")

	flag.StringVar(&screenShotPath, "screenshot-path", "data/screenshots", "Screenshot path")
	flag.Parse()

	config := Config{
		App: AppConfig{
			Listen: ListenHost{
				Host: listenHostFlag,
				Port: listenPortFlag,
			},
		},
		Postgres: Postgres{
			Host:     dbHostFlag,
			Port:     dbPortFlag,
			DB:       dbNameFlag,
			User:     dbUserNameFlag,
			Password: dbUserPasswordFlag,
		},
		Gowitness: Gowitness{
			ScreenshotPath: screenShotPath,
		},
	}
	// utils.D(config)
	return config
}
