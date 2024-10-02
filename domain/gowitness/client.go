package gowitness

import (
	"log/slog"

	"wsw/backend/ent"

	"github.com/sensepost/gowitness/pkg/runner"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	Client interface {
		UpdateUrl(*ent.Url)
	}
	CreateWriter func(*ent.Url) writers.Writer

	clientImlp struct {
		logger       *slog.Logger
		driver       runner.Driver
		options      runner.Options
		createWriter CreateWriter
	}
)

// UpdateUrl implements Client.
func (r clientImlp) UpdateUrl(uri *ent.Url) {
	r.logger.Info("Starting gowitness for: " + uri.URL)
	writer := r.createWriter(uri)
	runner := r.createRunner(writer)

	go func() {
		runner.Targets <- uri.URL
		close(runner.Targets)
	}()

	runner.Run()
	runner.Close()

	r.logger.Info("Finished gowitness for: " + uri.URL)
}

func (r clientImlp) createRunner(writer writers.Writer) runner.Runner {
	r.logger.Info("Creating gowitness runner")
	runner, err := runner.NewRunner(r.logger, r.driver, r.options, []writers.Writer{writer})
	if err != nil {
		r.logger.Error("Error creating runner", slog.Any("runner error", err))
	}
	return *runner
}

func NewClient(logger *slog.Logger, createWriter CreateWriter, driver runner.Driver, opts runner.Options) Client {
	return clientImlp{
		createWriter: createWriter,
		logger:       logger,
		driver:       driver,
		options:      opts,
	}
}
