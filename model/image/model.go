package image

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	Model interface {
		Insert(filename string, desctinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error)
		GetByID(int) (*ent.Image, error)
	}
	modelImpl struct {
		repository repository.ImageRepository
	}
)

// GetByID implements UploadImage.
func (u modelImpl) GetByID(imageID int) (*ent.Image, error) {
	return u.repository.GetByID(imageID)
}

// Insert implements UploadImage.
func (u modelImpl) Insert(filename string, desctinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error) {
	return u.repository.Insert(filename, desctinationPath, originalFilename, filetype, extraValue)
}

func NewModel(repository repository.ImageRepository) Model {
	return modelImpl{repository: repository}
}
