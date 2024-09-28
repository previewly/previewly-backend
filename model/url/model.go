package url

import (
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/preview"
	"wsw/backend/ent/repository"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		apiClient  gowitness.Client
		repository repository.Url
	}
)

// AddURL implements Url.
func (u urlImpl) AddURL(string) (*preview.PreviewData, error) {
	panic("unimplemented")
}

func NewUrl(urlRepository repository.Url, client gowitness.Client) Url {
	return urlImpl{repository: urlRepository, apiClient: client}
}
