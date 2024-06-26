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
		Status: convertPreviewStatus(data.Status),
	}
}

func convertPreviewStatus(status preview.Status) model.Status {
	switch status {
	case preview.StatusPending:
		return model.StatusPending
	case preview.StatusError:
		return model.StatusError
	case preview.StatusSuccess:
		return model.StatusSuccess
	default:
		return model.StatusPending
	}
}
