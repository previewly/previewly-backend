package gowitness

import (
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/domain/url"
	"wsw/backend/domain/url/screenshot"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	writerImpl struct {
		url                  *ent.Url
		repository           repository.Url
		urlProvider          screenshot.Provider
		relativePathProvider relative.Provider
	}
)

// Write implements writers.Writer.
func (w writerImpl) Write(result *models.Result) error {
	_, err := w.repository.Update(w.urlProvider.Provide(result.Filename), w.relativePathProvider.Provide(result.Filename), url.Success, w.url.ID, nil)
	return err
}

func NewRunnerWriter(url *ent.Url, repository repository.Url, urlProvider screenshot.Provider, relativePathProvider relative.Provider) writers.Writer {
	return writerImpl{
		url:                  url,
		repository:           repository,
		urlProvider:          urlProvider,
		relativePathProvider: relativePathProvider,
	}
}
