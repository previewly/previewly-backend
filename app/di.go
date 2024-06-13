package app

import (
	"context"
	"strings"
	"wsw/backend/ent"

	"github.com/golobby/container/v3"
)

func InitDi(config Config, appContext context.Context) {
	container.Singleton(func() context.Context { return appContext })
	container.Singleton(func() Config { return config })
	container.Singleton(func(config Config) (*ent.Client, error) {
		client, err := ent.Open("postgres", createConnectionURI(config.Postgres))
		if err != nil {
			return nil, err
		}
		if err := client.Schema.Create(appContext); err != nil {
			return nil, err
		}

		return client, nil
	})
}

func createConnectionURI(config Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     config.Port,
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.DB,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
}
