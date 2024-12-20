package process

import (
	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/graph/model"
)

type (
	Convertor interface {
		Convert(*ent.ImageProcess, *string, *string) *model.ImageProcess
	}
	convertorImpl struct{}
)

func NewConvertor() Convertor {
	return convertorImpl{}
}

func (c convertorImpl) Convert(processEntity *ent.ImageProcess, imageName *string, imageURL *string) *model.ImageProcess {
	return &model.ImageProcess{
		ID:        processEntity.ID,
		Image:     c.convertImageData(imageName, imageURL),
		Processes: c.convertToGQLProcesses(processEntity.Processes),
		Error:     &processEntity.Error,
		Status:    model.Status(processEntity.Status),
	}
}

func (c convertorImpl) convertImageData(imageName *string, imageURL *string) *model.ImageData {
	if imageURL != nil {
		return &model.ImageData{
			Name: *imageName,
			URL:  *imageURL,
		}
	}
	return nil
}

func (c convertorImpl) convertToGQLProcesses(processes []types.ImageProcess) []*model.OneImageProcess {
	gql := make([]*model.OneImageProcess, 0, len(processes))
	for _, process := range processes {
		gql = append(gql, &model.OneImageProcess{
			Type:    model.ImageProcessType(process.Type),
			Options: c.convertToGQLOptions(process.Options),
		})
	}
	return gql
}

func (c convertorImpl) convertToGQLOptions(options []types.ImageProcessOption) []*model.ImageProcessOption {
	gql := make([]*model.ImageProcessOption, 0, len(options))
	for _, option := range options {
		gql = append(gql, &model.ImageProcessOption{
			Key:   option.Key,
			Value: option.Value,
		})
	}
	return gql
}
