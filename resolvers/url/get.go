package url

import (
	"errors"

	"wsw/backend/graph/convertor"
	"wsw/backend/graph/model"
	tokenModel "wsw/backend/model/token"
	urlModel "wsw/backend/model/url"

	"github.com/golobby/container/v3"
)

func ResolveGetPreview(token string, url string) (*model.PreviewData, error) {
	var tokenModelImpl tokenModel.Token
	var urlModelImpl urlModel.Url

	errTokenModel := container.Resolve(&tokenModelImpl)
	errURLModel := container.Resolve(&urlModelImpl)

	if errURLModel != nil {
		return nil, errTokenModel
	}
	if errTokenModel != nil {
		return nil, errTokenModel
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}

	previewData, err := urlModelImpl.GetPreviewData(url)
	if err != nil {
		return nil, err
	}
	return convertor.ConvertPreviewData(previewData), nil
}
