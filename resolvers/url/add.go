package url

import (
	"errors"

	"wsw/backend/graph/convertor"
	"wsw/backend/graph/model"
	tokenModel "wsw/backend/model/token"
	urlModel "wsw/backend/model/url"

	"github.com/golobby/container/v3"
)

func ResolveAddURL(token string, url string) (*model.PreviewData, error) {
	var tokenModelImpl tokenModel.Token
	err := container.Resolve(&tokenModelImpl)
	if err != nil {
		return nil, err
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}
	var urlModelImpl urlModel.Url
	errURLModel := container.Resolve(&urlModelImpl)
	if errURLModel != nil {
		return nil, errURLModel
	}

	previewData, errData := urlModelImpl.AddURL(url)
	if errData != nil {
		return nil, errData
	}
	return convertor.ConvertPreviewData(previewData), nil
}
