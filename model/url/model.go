package url

import (
	"wsw/backend/domain/preview"
	"wsw/backend/ent/repository"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		repository repository.Url
	}
)

// AddURL implements Token.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	panic("unimplemented")
}

func NewUrl(urlRepository repository.Url) Url {
	return urlImpl{repository: urlRepository}
}
