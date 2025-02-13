package url

import (
	netUrl "net/url"
	"strings"

	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/image/url"
	"wsw/backend/domain/preview"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/ent/types"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
		GetPreviewData(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		client                gowitness.Client
		urlRepository         repository.Url
		screenshotURLProvider url.Provider
	}
)

// GetPreviewData implements Url.
func (u urlImpl) GetPreviewData(url string) (*preview.PreviewData, error) {
	entity, err := u.urlRepository.Get(url)
	if err != nil {
		return nil, err
	}

	return u.createPreviewData(entity)
}

// AddURL implements Url.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	_, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}
	urlEntity, err, isNew := u.getOrCreateURLEntity(url)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := u.updateURLData(urlEntity, isNew)
	if err != nil {
		return nil, err
	}

	return u.createPreviewData(updatedEntity)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client, provider url.Provider) Url {
	return urlImpl{
		urlRepository:         urlRepository,
		client:                client,
		screenshotURLProvider: provider,
	}
}

func (u urlImpl) updateURLData(urlEntity *ent.Url, isNew bool) (*ent.Url, error) {
	if isNew {
		go func(url *ent.Url) {
			u.client.UpdateUrl(url)
		}(urlEntity)
	}
	return urlEntity, nil
}

func (u urlImpl) getOrCreateURLEntity(url string) (*ent.Url, error, bool) {
	entity := u.urlRepository.TryGet(url)
	if entity == nil {
		newEntity, err := u.urlRepository.Insert(url)
		return newEntity, err, true
	}
	return entity, nil, false
}

func (u urlImpl) createPreviewData(url *ent.Url) (*preview.PreviewData, error) {
	lastError, err := u.getLastError(url)
	if err != nil {
		return nil, err
	}

	lastStat, err := u.getLastStat(url)
	if err != nil {
		return nil, err
	}

	var errorMessage *string = nil
	if lastError != nil {
		errorMessage = lastError.Message
	}

	var title *string = nil
	if lastStat != nil {
		title = lastStat.Title
	}
	imagePath := u.getImagePath(url.RelativePath)
	return &preview.PreviewData{
		ID:     url.ID,
		URL:    url.URL,
		Image:  u.screenshotURLProvider.Provide(imagePath),
		Status: u.getPreviewDataStatus(url.Status),
		Error:  errorMessage,
		Title:  title,
	}, nil
}

func (u urlImpl) getImagePath(path *string) *string {
	if path != nil {
		var sb strings.Builder
		sb.WriteString("screenshots/")
		sb.WriteString(*path)
		result := sb.String()
		return &result
	} else {
		return nil
	}
}

func (u urlImpl) getLastError(entity *ent.Url) (*ent.ErrorResult, error) {
	errors, error := u.urlRepository.GetErrors(entity)
	if error != nil {
		return nil, error
	}
	count := len(errors)
	if count == 0 {
		return nil, nil
	}
	return errors[count-1], nil
}

func (u urlImpl) getLastStat(entity *ent.Url) (*ent.Stat, error) {
	stats, error := u.urlRepository.GetStats(entity)
	if error != nil {
		return nil, error
	}
	count := len(stats)
	if count == 0 {
		return nil, nil
	}
	return stats[count-1], nil
}

func (u urlImpl) getPreviewDataStatus(status types.StatusEnum) preview.Status {
	switch status {
	case types.Success:
		return preview.StatusSuccess
	case types.Error:
		return preview.StatusError
	case types.Pending:
		return preview.StatusPending
	default:
		return preview.StatusPending
	}
}
