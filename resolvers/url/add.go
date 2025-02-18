package url

import (
	"errors"

	"wsw/backend/domain/gowitness"
	"wsw/backend/ent"
	"wsw/backend/graph/convertor"
	"wsw/backend/graph/model"
	tokenModel "wsw/backend/model/token"
	urlModel "wsw/backend/model/url"

	"github.com/golobby/container/v3"
)

func ResolveAddURL(token string, url string) (*model.PreviewData, error) {
	var tokenModelImpl tokenModel.Token
	var client gowitness.Client
	var urlModelImpl urlModel.Url

	err := container.Resolve(&tokenModelImpl)
	if err != nil {
		return nil, err
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}

	errURLModel := container.Resolve(&urlModelImpl)
	if errURLModel != nil {
		return nil, errURLModel
	}

	errClient := container.Resolve(&client)
	if errClient != nil {
		return nil, errClient
	}

	previewData, errData := urlModelImpl.AddURL(url)
	if errData != nil {
		return nil, errData
	}

	if previewData.IsNew {
		go func(url *ent.Url) {
			client.UpdateUrl(url)
		}(previewData.Entity)
	}

	return convertor.ConvertPreviewData(previewData), nil
}
