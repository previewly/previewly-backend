package gowitness

import (
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	writerImpl struct {
		url                  *ent.Url
		urlRepository        repository.Url
		statRepository       repository.Stat
		relativePathProvider relative.Provider
	}
)

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

func NewRunnerWriter(url *ent.Url, urlRepository repository.Url, statRepository repository.Stat, relativePathProvider relative.Provider) writers.Writer {
	return writerImpl{
		url:                  url,
		urlRepository:        urlRepository,
		statRepository:       statRepository,
		relativePathProvider: relativePathProvider,
	}
}
