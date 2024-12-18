package process

import (
	"context"

	"wsw/backend/graph/model"
	"wsw/backend/lib/utils"
	"wsw/backend/model/image"
)

type (
	Resolver interface {
		Resolve(context.Context, int, []*model.ImageProcessesInput) (*model.ImageProcesses, error)
	}

	resolverImpl struct {
		imagesModel image.UploadedImage
	}
)

func NewProcessResolver(imagesModel image.UploadedImage) Resolver {
	return resolverImpl{imagesModel: imagesModel}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(ctx context.Context, imageID int, processes []*model.ImageProcessesInput) (*model.ImageProcesses, error) {
	imageEntity, err := r.imagesModel.GetByID(imageID)
	if err != nil {
		return nil, err
	}
	utils.D(imageEntity)
	panic("unimplemented")
}
