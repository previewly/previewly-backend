package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"wsw/backend/app/config"
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/image/process"
	domainRunner "wsw/backend/domain/image/process"
	"wsw/backend/domain/image/upload"
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/domain/token/generator"
	"wsw/backend/domain/url/screenshot"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	log "wsw/backend/lib/log"
	"wsw/backend/lib/utils"
	"wsw/backend/model/token"
	"wsw/backend/model/url"

	domainImagePathProvider "wsw/backend/domain/image/path"
	domainStorage "wsw/backend/domain/image/upload/storage"

	imageModel "wsw/backend/model/image"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golobby/container/v3"
	"github.com/rollbar/rollbar-go"
	"github.com/rs/cors"
	"github.com/sensepost/gowitness/pkg/runner"
	driver "github.com/sensepost/gowitness/pkg/runner/drivers"
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

	initService(func() *slog.Logger { return slog.Default() })

	initService(func(urlRepository repository.Url, statRepository repository.Stat, relativePathProvider relative.Provider) gowitness.CreateWriter {
		return func(url *ent.Url) gowitness.Writer {
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

	initModels()
	initDomains(config)
	initResolvers()
}

func initDomains(config config.Config) {
	initService(func() domainStorage.FilenameGenerator { return domainStorage.NewFilenameProvider() })
	initService(func() domainImagePathProvider.PathProvider {
		return domainImagePathProvider.NewPathProvider(config.App.UploadPath)
	})
	initService(func(filenameGenerator domainStorage.FilenameGenerator, pathProvider domainImagePathProvider.PathProvider) domainStorage.Storage {
		return domainStorage.NewUploadStorage(filenameGenerator, pathProvider)
	})
	initService(func() process.Convertor { return process.NewConvertor() })
	initService(func() domainRunner.ProcessFactory { return domainRunner.NewProcessFactory() })
	initService(func(pathProvider domainImagePathProvider.PathProvider, processFactory domainRunner.ProcessFactory) domainRunner.ProcessRunner {
		return domainRunner.NewProcessRunner(pathProvider, processFactory)
	})
}

func initResolvers() {
	initService(func(model imageModel.UploadedImage, storage domainStorage.Storage) upload.Resolver {
		return upload.NewUploadResolver(model, storage)
	})
	initService(func(model imageModel.UploadedImage, processesModel imageModel.ImageProcesses, gqlConvertor process.Convertor, runner domainRunner.ProcessRunner) process.Resolver {
		return process.NewProcessResolver(model, processesModel, gqlConvertor, runner)
	})
}

func initModels() {
	initService(func(generator generator.TokenGenerator, tokenRepository repository.Token) token.Token {
		return token.NewModel(generator, tokenRepository)
	})
	initService(func(urlRepository repository.Url, client gowitness.Client, provider screenshot.Provider) url.Url {
		return url.NewUrl(urlRepository, client, provider)
	})

	initService(func(uploadRepository repository.UploadImageRepository) imageModel.UploadedImage {
		return imageModel.NewModel(uploadRepository)
	})
	initService(func(processRepository repository.ImageProcessRepository, imageRepository repository.UploadImageRepository) imageModel.ImageProcesses {
		return imageModel.NewImageProcesses(processRepository, imageRepository)
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
	initService(func(client *ent.Client, ctx context.Context) repository.UploadImageRepository {
		return repository.NewUploadImageRepository(client, ctx)
	})
	initService(func(client *ent.Client, ctx context.Context) repository.ImageProcessRepository {
		return repository.NewImageProcessRepository(client, ctx)
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
