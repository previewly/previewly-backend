package url

import (
	netUrl "net/url"
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/preview"
	"wsw/backend/domain/url"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/lib/utils"
)

type (
	Url interface {
		AddURL(string) (*preview.PreviewData, error)
	}
	urlImpl struct {
		apiClient  gowitness.Client
		repository repository.Url
	}
	UrlEntityHolder struct {
		entity *ent.Url
		isNew  bool
	}
)

func (u urlImpl) getUrlEntity(url string) (*UrlEntityHolder, error) {
	urlEntity := u.repository.TryGet(url)
	if urlEntity == nil {
		urlEntity, err := u.repository.Insert(url)
		if err != nil {
			return nil, err
		}
		return &UrlEntityHolder{entity: urlEntity, isNew: true}, nil
	}
	return &UrlEntityHolder{entity: urlEntity, isNew: false}, nil
}

func (u urlImpl) updateUrlData(url *ent.Url, isNew bool) error {
	if u.shouldAddUrlToApi(url, isNew) {
		go func(url *ent.Url) {
			id, err := u.apiClient.AddURL(url.URL)
			u.setApiUrlId(url, id, err)
		}(url)
	}
	details, err := u.apiClient.Details(url)
	if err != nil {
		return err
	}
	u.updateApiURLDetails(details)
	return nil
}

// AddURL implements Token.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	_, err := netUrl.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}

	urlEntityHolder, err := u.getUrlEntity(url)
	if err != nil {
		return nil, err
	}
	errUpdate := u.updateUrlData(urlEntityHolder.entity, urlEntityHolder.isNew)
	if errUpdate != nil {
		return nil, errUpdate
	}
	preview, errPreview := u.getPreviewData(urlEntityHolder.entity)
	if errPreview != nil {
		return nil, errPreview
	}
	return preview, nil
}

func (u urlImpl) getPreviewData(url *ent.Url) (*preview.PreviewData, error) {
	return &preview.PreviewData{
		ID:     url.ID,
		URL:    url.URL,
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

func (u urlImpl) shouldAddUrlToApi(_ *ent.Url, isNew bool) bool {
	// TODO
	return isNew
}

func (u urlImpl) updateApiURLDetails(details gowitness.DetailsURL) {
	u.repository.Update(details.Image, details.ID)
}

func (u urlImpl) setApiUrlId(url *ent.Url, apiUrlId int, urlError error) {
	err := u.repository.UpdateApiUrlId(url, apiUrlId)
	utils.D(err)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client) Url {
	return urlImpl{repository: urlRepository, apiClient: client}
}
