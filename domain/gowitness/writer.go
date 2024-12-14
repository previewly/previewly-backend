package gowitness

import (
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	Writer interface {
		writers.Writer
		Error(error) error
	}
	writerImpl struct {
		url                  *ent.Url
		urlRepository        repository.Url
		statRepository       repository.Stat
		relativePathProvider relative.Provider
	}
)

// Error implements Writer.
func (w writerImpl) Error(err error) error {
	_, saveErr := w.urlRepository.SaveFailure(err.Error(), w.url.ID)
	return saveErr
}

// Write implements writers.Writer.
func (w writerImpl) Write(result *models.Result) error {
	statEntity, errStat := w.statRepository.Insert(&result.Title)
	if errStat != nil {
		return errStat
	}
	_, err := w.urlRepository.SaveSuccess(w.relativePathProvider.Provide(result.Filename), statEntity, w.url.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewRunnerWriter(url *ent.Url, urlRepository repository.Url, statRepository repository.Stat, relativePathProvider relative.Provider) Writer {
	return writerImpl{
		url:                  url,
		urlRepository:        urlRepository,
		statRepository:       statRepository,
		relativePathProvider: relativePathProvider,
	}
}
