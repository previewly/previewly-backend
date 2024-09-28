package main

import (
	"log/slog"

	"wsw/backend/lib/utils"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/runner"
	driver "github.com/sensepost/gowitness/pkg/runner/drivers"
	"github.com/sensepost/gowitness/pkg/writers"
)

type customWriter struct {
	Result *models.Result
}

// Write results to stdout
func (c *customWriter) Write(result *models.Result) error {
	c.Result = result
	return nil
}

func NewCustomWriter() (*customWriter, error) {
	return &customWriter{}, nil
}

func main() {
	logger := slog.Default()

	// define scan/chrome/logging etc. options. drivers, scanners and writers use these.
	// it includes concurrency options, where to save screenshots and more.
	opts := runner.NewDefaultOptions()
	// set any opts you want, or start from scratch with &runner.Options{}

	// get the driver, the writer and a runner that glues it all together
	driver, _ := driver.NewChromedp(logger, *opts)
	writer, _ := NewCustomWriter()
	runner, _ := runner.NewRunner(logger, driver, *opts, []writers.Writer{writer})

	// with the runner up, you have runner.Targets which is a channel you can write targets to.
	// write from a goroutine, as the runner's Run()  method will wait for the `Target` channel to close.

	// target "pusher" goroutine
	go func() {
		runner.Targets <- "https://sensepost.com"
		close(runner.Targets)
	}()

	// finally, run the runner, when done, close it
	runner.Run() // will block until runner.Targets is closed
	utils.D(writer.Result)
	runner.Close()
}
