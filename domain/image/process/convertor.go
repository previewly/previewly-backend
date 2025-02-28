package process

import (
	"wsw/backend/domain/image/process/runner/result"
	"wsw/backend/ent/types"
	"wsw/backend/graph/model"

	"github.com/xorcare/pointer"
)

type (
	Convertor interface {
		Convert(result.Result) *model.ImageProcess
	}
	convertorImpl struct{}
)

func NewConvertor() Convertor {
	return convertorImpl{}
}

func (c convertorImpl) Convert(result result.Result) *model.ImageProcess {
	var errorMessage *string
	if result.Error != nil {
		errorMessage = pointer.String(result.Error.Error())
	}

	return &model.ImageProcess{
		Image:     c.convertImageData(result.Input.Image.Filename, result.ImageURL),
		Processes: c.convertToGQLProcesses(result.Input.Processes),
		Error:     errorMessage,
		Status:    model.Status(result.Status),
	}
}

func (c convertorImpl) convertImageData(imageName string, imageURL *string) *model.ImageData {
	if imageURL != nil {
		return &model.ImageData{
			Name: imageName,
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
