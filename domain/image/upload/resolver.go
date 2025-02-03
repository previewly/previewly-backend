package upload

import (
	"context"
	"errors"

	"wsw/backend/domain/image"
	"wsw/backend/ent"
	"wsw/backend/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/xorcare/pointer"
)

const MaxImageSize = 5 * 1024 * 1024

type (
	Resolver interface {
		Resolve(context.Context, []*model.UploadInput) ([]*model.UploadImageStatus, error)
	}
	resolverImpl struct {
		saver image.Saver
	}
)

func NewUploadResolver(saver image.Saver) Resolver {
	return resolverImpl{saver: saver}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(cxt context.Context, images []*model.UploadInput) ([]*model.UploadImageStatus, error) {
	fileResults := make([]*model.UploadImageStatus, len(images))
	for i, image := range images {
		if image == nil {
			return nil, errors.New("image is nil")
		}

		imageError := r.validateImage(&image.Image)
		if imageError != nil {
			fileResults[i] = r.createErrorStatus(image.Image.Filename, imageError, image.Extra)
		} else {
			imageEntity, err := r.saver.SaveImage(image.Image.Filename, image.Image.File, image.Image.ContentType, image.Extra)

			if err != nil {
				fileResults[i] = r.createErrorStatus(image.Image.Filename, imageError, image.Extra)
			} else {
				fileResults[i] = r.createImageStatus(image.Image.Filename, *imageEntity)
			}
		}
	}
	return fileResults, nil
}

func (r resolverImpl) validateImage(image *graphql.Upload) error {
	if image.Size == 0 {
		return errors.New("em")
	}
	if image.Size > MaxImageSize {
		return errors.New("image is too big")
	}
	if image.ContentType != "image/jpeg" && image.ContentType != "image/jpg" && image.ContentType != "image/png" {
		return errors.New("unsupported image format")
	}
	return nil
}

func (r resolverImpl) createErrorStatus(name string, imageError error, extra *string) *model.UploadImageStatus {
	return &model.UploadImageStatus{
		Name:   name,
		Error:  pointer.String(imageError.Error()),
		Status: model.StatusError,
		Extra:  extra,
	}
}

func (r resolverImpl) createImageStatus(name string, imageEntity ent.UploadImage) *model.UploadImageStatus {
	return &model.UploadImageStatus{
		ID:     imageEntity.ID,
		Name:   name,
		Error:  nil,
		Status: model.StatusSuccess,
		Extra:  imageEntity.ExtraValue,
	}
}
