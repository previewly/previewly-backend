package url

import (
	netUrl "net/url"

	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/preview"
	"wsw/backend/domain/url"
	"wsw/backend/domain/url/screenshot"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		client                gowitness.Client
		repository            repository.Url
		screenshotURLProvider screenshot.Provider
	}
)

// AddURL implements Url.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	_, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}
	urlEntity, err, isNew := u.getUrlEntity(url)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := u.updateUrlData(urlEntity, isNew)
	if err != nil {
		return nil, err
	}
	return u.getPreviewData(updatedEntity)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client, provider screenshot.Provider) Url {
	return urlImpl{repository: urlRepository, client: client, screenshotURLProvider: provider}
}

func (u urlImpl) updateUrlData(urlEntity *ent.Url, isNew bool) (*ent.Url, error) {
	if isNew {
		go func(url *ent.Url) {
			u.client.UpdateUrl(url)
		}(urlEntity)
	}
	return urlEntity, nil
}

func (u urlImpl) getUrlEntity(url string) (*ent.Url, error, bool) {
	entity := u.repository.TryGet(url)
	if entity == nil {
		entity, err := u.repository.Insert(url)
		return entity, err, true
	}
	return entity, nil, false
}

func (u urlImpl) getPreviewData(url *ent.Url) (*preview.PreviewData, error) {
	return &preview.PreviewData{
		ID:     url.ID,
		URL:    url.URL,
		Image:  u.screenshotURLProvider.Provide(url.RelativePath),
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
