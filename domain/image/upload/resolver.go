package upload

import (
	"context"
	"errors"

	"wsw/backend/domain/image/upload/storage"
	"wsw/backend/ent"
	"wsw/backend/graph/model"
	"wsw/backend/model/image"

	"github.com/99designs/gqlgen/graphql"
	"github.com/xorcare/pointer"
)

const MaxImageSize = 5 * 1024 * 1024

type (
	Resolver interface {
		Resolve(context.Context, []*model.UploadInput) ([]*model.UploadImageStatus, error)
	}
	resolverImpl struct {
		model   image.UploadedImage
		storage storage.Storage
	}
)

func NewUploadResolver(model image.UploadedImage, storage storage.Storage) Resolver {
	return resolverImpl{model: model, storage: storage}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(cxt context.Context, images []*model.UploadInput) ([]*model.UploadImageStatus, error) {
	fileResults := make([]*model.UploadImageStatus, len(images))
	for i, image := range images {
		if image == nil {
			return nil, errors.New("image is nil")
		}

		imageError := r.validateImage(&image.Image)
		storageFile, imageError := r.saveToStorage(&image.Image, imageError)
		imageEntity, imageError := r.saveToDatabase(&image.Image, storageFile.NewFilePlace, storageFile.NewFilename, imageError, image.Extra)

		fileResults[i] = r.createImageStatus(image.Image.Filename, imageError, imageEntity)
	}
	return fileResults, nil
}

func (r resolverImpl) saveToStorage(image *graphql.Upload, imageError error) (*storage.StorageFile, error) {
	if imageError != nil {
		return nil, imageError
	}
	return r.storage.Save(image.Filename, pointer.String("o/"), image.File)
}

func (r resolverImpl) saveToDatabase(image *graphql.Upload, destinationPath string, newFilename string, imageError error, extraValue *string) (*ent.UploadImage, error) {
	if imageError != nil {
		return nil, imageError
	}

	return r.model.Insert(newFilename, destinationPath, image.Filename, image.ContentType, extraValue)
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

func (r resolverImpl) createImageStatus(name string, imageError error, imageEntity *ent.UploadImage) *model.UploadImageStatus {
	if imageError != nil {
		return &model.UploadImageStatus{
			Name:   name,
			Error:  pointer.String(imageError.Error()),
			Status: model.StatusError,
			Extra:  imageEntity.ExtraValue,
		}
	} else {
		return &model.UploadImageStatus{
			ID:     imageEntity.ID,
			Name:   name,
			Error:  nil,
			Status: model.StatusSuccess,
		}
	}
}
