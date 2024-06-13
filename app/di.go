package app

import (
	"context"
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"

	"github.com/golobby/container/v3"
)

func initDi(config Config, appContext context.Context) {
	initService(func() context.Context { return appContext })
	initService(func() Config { return config })
	initService(func(config Config) (*ent.Client, error) { return newDBClient(config.Postgres, appContext) })
	initService(func() App { return appImpl{router: newRouter()} })

	initService(func() generator.TokenGenerator { return generator.NewTokenGenerator() })
	initService(func() repository.Token { return repository.NewToken() })
	initService(func(generator generator.TokenGenerator) token.Token { return token.NewModel(generator) })
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
