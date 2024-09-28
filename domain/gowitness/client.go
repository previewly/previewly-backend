package gowitness

import (
	"log/slog"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/runner"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	Client interface {
		GetUrlDetails(string) (*Details, error)
	}
	Runner interface {
		Close()
	}
	Writer interface {
		Write(*models.Result) error
	}
	Details    struct{}
	clientImpl struct{}
	runnerImpl struct {
		runner *runner.Runner
	}
	writerImpl struct {
		Result *models.Result
	}
)

// Write implements Writer.
func (w writerImpl) Write(result *models.Result) error {
	w.Result = result
	return nil
}

// Close implements Runner.
func (r runnerImpl) Close() {
	panic("unimplemented")
}

func NewClient() Client {
	return clientImpl{}
}

func NewRunner(logger *slog.Logger, writer Writer, driver runner.Driver, opts runner.Options) Runner {
	runner, _ := runner.NewRunner(logger, driver, opts, []writers.Writer{writer})

	return runnerImpl{
		runner: runner,
	}
}

func NewRunnerWriter() Writer {
	return writerImpl{}
}

// GetUrlDetails implements Client.
func (c clientImpl) GetUrlDetails(string) (*Details, error) {
	panic("unimplemented")
}
