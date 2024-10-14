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
		GetPreviewData(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		client                gowitness.Client
		urlRepository         repository.Url
		screenshotURLProvider screenshot.Provider
	}
)

// GetPreviewData implements Url.
func (u urlImpl) GetPreviewData(url string) (*preview.PreviewData, error) {
	entity, err := u.urlRepository.Get(url)
	if err != nil {
		return nil, err
	}

	lastError, err := u.getLastError(entity)
	if err != nil {
		return nil, err
	}
	return u.getPreviewData(entity, lastError)
}

// AddURL implements Url.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	_, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}
	urlEntity, err, isNew := u.getOrCreateUrlEntity(url)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := u.updateUrlData(urlEntity, isNew)
	if err != nil {
		return nil, err
	}

	lastError, err := u.getLastError(updatedEntity)
	if err != nil {
		return nil, err
	}
	return u.getPreviewData(updatedEntity, lastError)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client, provider screenshot.Provider) Url {
	return urlImpl{
		urlRepository:         urlRepository,
		client:                client,
		screenshotURLProvider: provider,
	}
}

func (u urlImpl) updateUrlData(urlEntity *ent.Url, isNew bool) (*ent.Url, error) {
	if isNew {
		go func(url *ent.Url) {
			u.client.UpdateUrl(url)
		}(urlEntity)
	}
	return urlEntity, nil
}

func (u urlImpl) getOrCreateUrlEntity(url string) (*ent.Url, error, bool) {
	entity := u.urlRepository.TryGet(url)
	if entity == nil {
		entity, err := u.urlRepository.Insert(url)
		return entity, err, true
	}
	return entity, nil, false
}

func (u urlImpl) getPreviewData(url *ent.Url, lastError *ent.ErrorResult) (*preview.PreviewData, error) {
	return &preview.PreviewData{
		ID:     url.ID,
		URL:    url.URL,
		Image:  u.screenshotURLProvider.Provide(url.RelativePath),
		Status: u.getPreviewDataStatus(url.Status),
		Error:  lastError.Message,
	}, nil
}

func (u urlImpl) getLastError(entity *ent.Url) (*ent.ErrorResult, error) {
	errors, error := u.urlRepository.GetErrors(entity)
	if error != nil {
		return nil, error
	}
	return errors[len(errors)-1], nil
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
