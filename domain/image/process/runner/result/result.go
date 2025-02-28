package result

import (
	"errors"

	"wsw/backend/domain/dto"
	"wsw/backend/domain/image/process/input"
	"wsw/backend/domain/image/url"
	"wsw/backend/ent"
	"wsw/backend/ent/types"
)

type (
	Result struct {
		Input      input.Input
		PrefixPath string
		ImageURL   *string
		Status     types.StatusEnum
		Error      error
	}
	Factory interface {
		New(input input.Input, processEntity *ent.ImageProcess) (*Result, error)
	}
	factoryImpl struct {
		urlProvider url.Provider
	}
)

func NewFactory(urlProvider url.Provider) Factory { return factoryImpl{urlProvider: urlProvider} }

// New implements Factory.
func (f factoryImpl) New(input input.Input, processEntity *ent.ImageProcess) (*Result, error) {
	url := f.urlProvider.Provide(dto.NewImage(input.Image.Filename, processEntity.PathPrefix))

	var errorResult error
	if processEntity.Error != "" {
		errorResult = errors.New(processEntity.Error)
	}

	return &Result{
		Input:      input,
		PrefixPath: processEntity.PathPrefix,
		ImageURL:   &url,
		Status:     processEntity.Status,
		Error:      errorResult,
	}, nil
}
