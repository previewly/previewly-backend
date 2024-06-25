package convertor

import (
	"wsw/backend/domain/preview"
	"wsw/backend/graph/model"
)

func ConvertPreviewData(data *preview.PreviewData) *model.PreviewData {
	return &model.PreviewData{
		ID:     data.ID,
		URL:    data.URL,
		Image:  nil,
		Status: model.StatusSuccess,
	}
}
