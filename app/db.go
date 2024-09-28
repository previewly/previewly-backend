package app

import (
	"context"
	"strconv"
	"strings"

	"wsw/backend/app/config"
	"wsw/backend/ent"
	"wsw/backend/ent/migrate"

	_ "github.com/lib/pq"
)

func newDBClient(config config.Postgres, appContext context.Context) (*ent.Client, error) {
	client, err := ent.Open("postgres", createConnectionURI(config))
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		appContext,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		return nil, err
	}

	return client, nil
}

func createConnectionURI(config config.Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     strconv.Itoa(config.Port),
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
