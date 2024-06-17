package url

import (
	"wsw/backend/domain/preview"
	"wsw/backend/ent"
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

func (u urlImpl) getUrlEntity(url string) (*ent.Url, error) {
	urlEntity := u.repository.TryGet(url)
	if urlEntity == nil {
		urlEntity, err := u.repository.Insert(url)
		if err != nil {
			return nil, err
		}
		return urlEntity, nil
	}
	return urlEntity, nil
}

func (u urlImpl) updateUrlData(erl *ent.Url) (*preview.PreviewData, error) {
	panic("not implemented")
}

// AddURL implements Token.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	urlEntity, err := u.getUrlEntity(url)
	if err != nil {
		return nil, err
	}
	preview, errPreview := u.updateUrlData(urlEntity)
	if errPreview != nil {
		return nil, errPreview
	}
	return preview, nil
}

func NewUrl(urlRepository repository.Url) Url {
	return urlImpl{repository: urlRepository}
}
