package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"wsw/backend/app/config"
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/domain/token/generator"
	"wsw/backend/domain/url/screenshot"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	log "wsw/backend/lib/log"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"
	"wsw/backend/model/url"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golobby/container/v3"
	"github.com/rollbar/rollbar-go"
	"github.com/rs/cors"
	"github.com/sensepost/gowitness/pkg/runner"
	driver "github.com/sensepost/gowitness/pkg/runner/drivers"
	writers "github.com/sensepost/gowitness/pkg/writers"
)

type (
	Middlewares struct {
		List []func(http.Handler) http.Handler
	}
)

func initDi(config config.Config, appContext context.Context) {
	initVoidServices(config)

	initService(func() context.Context { return appContext })
	initService(func() (*ent.Client, error) { return newDBClient(config.Postgres, appContext) })

	initService(func() Middlewares {
		return Middlewares{
			List: []func(http.Handler) http.Handler{
				middleware.Logger,
				middleware.Recoverer,
				middleware.RealIP,
				cors.New(cors.Options{
					AllowedOrigins:     []string{"*"},
					AllowCredentials:   true,
					AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
					AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
					OptionsPassthrough: true,
					Debug:              true,
				}).
					Handler,
				sentryhttp.New(sentryhttp.Options{
					Repanic: true,
				}).Handle,
			},
		}
	})

	initService(func(entClient *ent.Client, middlewares Middlewares) App {
		return appImpl{
			router: newRouter(middlewares),
			listen: config.App.Listen,
			closer: func() {
				entClient.Close()
				sentry.Flush(time.Second)
				rollbar.Close()
			},
		}
	})

	initService(func() generator.TokenGenerator { return generator.NewTokenGenerator() })
	initService(func() screenshot.Loader { return screenshot.NewLoader(config.App.AssetsBaseURL) })
	initService(func(loader screenshot.Loader) screenshot.Provider {
		return screenshot.NewProvider(config.Gowitness, loader)
	})
	initService(func() relative.Provider { return relative.NewProvider() })

	initRepositories()

	initService(func() *slog.Logger {
		return slog.Default()
	})

	initService(func(urlRepository repository.Url, statRepository repository.Stat, relativePathProvider relative.Provider) gowitness.CreateWriter {
		return func(url *ent.Url) writers.Writer {
			return gowitness.NewRunnerWriter(url, urlRepository, statRepository, relativePathProvider)
		}
	})
	initService(func(logger *slog.Logger, createWriter gowitness.CreateWriter) gowitness.Client {
		logger.Info("Starting gowitness")

		options := runner.NewDefaultOptions()
		options.Scan.ScreenshotPath = config.Gowitness.ScreenshotPath

		driver, err := driver.NewChromedp(logger, *options)
		if err != nil {
			utils.F("Could not create driver: %v", err)
		}
		return gowitness.NewClient(logger, createWriter, driver, *options)
	})

	initService(func(generator generator.TokenGenerator, tokenRepository repository.Token) token.Token {
		return token.NewModel(generator, tokenRepository)
	})
	initService(func(urlRepository repository.Url, client gowitness.Client, provider screenshot.Provider) url.Url {
		return url.NewUrl(urlRepository, client, provider)
	})
}

func initRepositories() {
	initService(func(client *ent.Client, ctx context.Context) repository.Token {
		return repository.NewToken(client, ctx)
	})
	initService(func(client *ent.Client, ctx context.Context) repository.Url {
		return repository.NewUrl(client, ctx)
	})
	initService(func(client *ent.Client, ctx context.Context) repository.Stat {
		return repository.NewStat(client, ctx)
	})
}

func initVoidServices(config config.Config) {
	initSentry(config.Sentry)
	initRollbar(config.Rollbar)
	initLogger()
}

func initRollbar(config config.Rollbar) {
	rollbar.SetToken(config.Token)
}

func initSentry(options sentry.ClientOptions) {
	err := sentry.Init(options)
	if err != nil {
		utils.F("Sentry initialization failed: %v\n", err)
	}
}

func initLogger() {
	logger := slog.New(log.NewHandler(nil))
	slog.SetDefault(logger)
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
