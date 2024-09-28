package app

import (
	"context"
	"log/slog"

	"wsw/backend/app/config"
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/token/generator"
	"wsw/backend/domain/url/screenshot"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"
	"wsw/backend/model/url"

	"github.com/golobby/container/v3"
	"github.com/sensepost/gowitness/pkg/runner"
	driver "github.com/sensepost/gowitness/pkg/runner/drivers"

	writers "github.com/sensepost/gowitness/pkg/writers"
)

func initDi(config config.Config, appContext context.Context) {
	initService(func() context.Context { return appContext })
	initService(func() (*ent.Client, error) { return newDBClient(config.Postgres, appContext) })

	initService(func(entClient *ent.Client) App {
		return appImpl{
			router: newRouter(),
			listen: config.App.Listen,
			closer: func() {
				entClient.Close()
			},
		}
	})

	initService(func() generator.TokenGenerator { return generator.NewTokenGenerator() })
	initService(func() screenshot.Provider { return screenshot.NewProvider(config.Gowitness) })

	initService(func(client *ent.Client, ctx context.Context) repository.Token {
		return repository.NewToken(client, ctx)
	})
	initService(func(client *ent.Client, ctx context.Context) repository.Url {
		return repository.NewUrl(client, ctx)
	})

	initService(func() *slog.Logger { return slog.Default() })
	initService(func(repository repository.Url, urlProvider screenshot.Provider) gowitness.CreateWriter {
		return func(url *ent.Url) writers.Writer {
			return gowitness.NewRunnerWriter(url, repository, urlProvider)
		}
	})
	initService(func(logger *slog.Logger, createWriter gowitness.CreateWriter) gowitness.Client {
		options := runner.NewDefaultOptions()
		options.Scan.ScreenshotPath = config.Gowitness.ScreenshotPath

		driver, _ := driver.NewChromedp(logger, *options)
		return gowitness.NewClient(logger, createWriter, driver, *options)
	})

	initService(func(generator generator.TokenGenerator, tokenRepository repository.Token) token.Token {
		return token.NewModel(generator, tokenRepository)
	})
	initService(func(urlRepository repository.Url, client gowitness.Client) url.Url {
		return url.NewUrl(urlRepository, client)
	})
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
