package upload

import (
	"context"
	"errors"

	"wsw/backend/graph/model"
	"wsw/backend/model/upload"

	"github.com/99designs/gqlgen/graphql"
	"github.com/xorcare/pointer"
)

const MaxImageSize = 5 * 1024 * 1024

type (
	Resolver interface {
		Resolve(context.Context, []*graphql.Upload) ([]*model.UploadImageStatus, error)
	}
	resolverImpl struct {
		model upload.UploadImage
	}
)

func NewUploadResolver(model upload.UploadImage) Resolver {
	return resolverImpl{model: model}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(cxt context.Context, images []*graphql.Upload) ([]*model.UploadImageStatus, error) {
	fileResults := make([]*model.UploadImageStatus, len(images))
	for i, image := range images {
		if image == nil {
			return nil, errors.New("image is nil")
		}

		imageStatus := r.createImageStatus(image.Filename)
		imageError := r.validateImage(image)
		if imageError == nil {
			imageEntity, err := r.model.Insert(image.Filename, image.ContentType)
			if err != nil {
				imageError = pointer.String(err.Error())
			} else {
				imageStatus.ID = imageEntity.ID
			}
		}
		if imageError == nil {
			imageStatus.Error = nil
			imageStatus.Status = model.StatusSuccess
		} else {
			imageStatus.Error = imageError
			imageStatus.Status = model.StatusError
		}

		fileResults[i] = &imageStatus
	}
	return fileResults, nil
}

func (r resolverImpl) validateImage(image *graphql.Upload) *string {
	if image.Size == 0 {
		return pointer.String("empty image")
	}
	if image.Size > MaxImageSize {
		return pointer.String("image is too big")
	}
	if image.ContentType != "image/jpeg" && image.ContentType != "image/jpg" && image.ContentType != "image/png" {
		return pointer.String("unsupported image format")
	}
	return nil
}

func (r resolverImpl) createImageStatus(name string) model.UploadImageStatus {
	return model.UploadImageStatus{
		Name:   name,
		Error:  nil,
		Status: model.StatusPending,
	}
}
