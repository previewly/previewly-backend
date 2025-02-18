package image

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	UploadedImage interface {
		Insert(filename string, desctinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error)
		GetByID(int) (*ent.Image, error)
	}
	uploadImpl struct {
		repository repository.ImageRepository
	}
)

// GetByID implements UploadImage.
func (u uploadImpl) GetByID(imageID int) (*ent.Image, error) {
	return u.repository.GetByID(imageID)
}

// Insert implements UploadImage.
func (u uploadImpl) Insert(filename string, desctinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error) {
	return u.repository.Insert(filename, desctinationPath, originalFilename, filetype, extraValue)
}

func NewModel(repository repository.ImageRepository) UploadedImage {
	return uploadImpl{repository: repository}
}
