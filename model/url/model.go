package url

import (
	netUrl "net/url"

	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/preview"
	"wsw/backend/domain/url"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		client     gowitness.Client
		repository repository.Url
	}
)

// AddURL implements Url.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	_, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}
	urlEntity, err := u.getUrlEntity(url)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := u.updateUrlData(urlEntity)
	if err != nil {
		return nil, err
	}
	return u.getPreviewData(updatedEntity)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client) Url {
	return urlImpl{repository: urlRepository, client: client}
}

func (u urlImpl) updateUrlData(urlEntity *ent.Url) (*ent.Url, error) {
	if u.shouldUpdateUrlData(urlEntity) {
		go func(url *ent.Url) {
			u.client.UpdateUrl(url)
		}(urlEntity)
	}
	return urlEntity, nil
}

func (u urlImpl) shouldUpdateUrlData(urlEntity *ent.Url) bool {
	if urlEntity.Status == url.Success || urlEntity.Status == url.Error {
		return false
	}
	return true
}

func (u urlImpl) getUrlEntity(url string) (*ent.Url, error) {
	entity := u.repository.TryGet(url)
	if entity == nil {
		return u.repository.Insert(url)
	}
	return entity, nil
}

func (u urlImpl) getPreviewData(url *ent.Url) (*preview.PreviewData, error) {
	return &preview.PreviewData{
		ID:     url.ID,
		URL:    url.URL,
		Image:  url.ImageURL,
		Status: u.getPreviewDataStatus(url.Status),
	}, nil
}

func (u urlImpl) getPreviewDataStatus(status url.Status) preview.Status {
	switch status {
	case url.Success:
		return preview.StatusSuccess
	case url.Error:
		return preview.StatusError
	case url.Pending:
		return preview.StatusPending
	default:
		return preview.StatusPending
	}
}
