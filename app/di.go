package app

import (
	"context"
	"wsw/backend/ent"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"

	"github.com/golobby/container/v3"
)

func initDi(config Config, appContext context.Context) {
	initService(func() context.Context { return appContext })
	initService(func() Config { return config })
	initService(func(config Config) (*ent.Client, error) { return newDBClient(config.Postgres, appContext) })
	initService(func() App { return appImpl{router: newRouter()} })
	initService(func() token.Token { return token.NewModel() })
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
