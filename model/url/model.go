package url

import (
	"wsw/backend/domain/gowitness"
	"wsw/backend/domain/preview"
	"wsw/backend/ent"
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
			id, err := u.apiClient.AddUrl(url.URL)
			u.setApiUrlId(url, id, err)
		}(url)
	}
	if url.APIURLID != nil {
		details, err := u.apiClient.Details(*url.APIURLID)
		if err != nil {
			return err
		}
		u.updateApiURLDetails(details)
	}
	return nil
}

// AddURL implements Token.
func (u urlImpl) AddURL(url string) (*preview.PreviewData, error) {
	urlEntityHolder, err := u.getUrlEntity(url)
	if err != nil {
		return nil, err
	}
	u.updateUrlData(urlEntityHolder.entity, urlEntityHolder.isNew)

	preview, errPreview := u.getPreviewData(urlEntityHolder.entity)
	if errPreview != nil {
		return nil, errPreview
	}
	return preview, nil
}

func (u urlImpl) getPreviewData(url *ent.Url) (*preview.PreviewData, error) {
	return &preview.PreviewData{
		ID:  url.ID,
		URL: url.URL,
	}, nil
}

func (u urlImpl) shouldAddUrlToApi(url *ent.Url, isNew bool) bool {
	// TODO
	return isNew
}

func (u urlImpl) updateApiURLDetails(details gowitness.DetailsURL) {
	panic("updateApiURLDetails not implemented")
}

func (u urlImpl) setApiUrlId(url *ent.Url, apiUrlId int, urlError error) {
	u.repository.UpdateApiUrlId(url, apiUrlId)
}

func NewUrl(urlRepository repository.Url, client gowitness.Client) Url {
	return urlImpl{repository: urlRepository, apiClient: client}
}
