package image

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	UploadedImage interface {
		Insert(string, string, string, string) (*ent.UploadImage, error)
		GetByID(int) (*ent.UploadImage, error)
	}
	uploadImpl struct {
		repository repository.UploadImageRepository
	}
)

// GetByID implements UploadImage.
func (u uploadImpl) GetByID(imageID int) (*ent.UploadImage, error) {
	return u.repository.GetByID(imageID)
}

// Insert implements UploadImage.
func (u uploadImpl) Insert(filename string, desctinationPath string, originalFilename string, filetype string) (*ent.UploadImage, error) {
	return u.repository.Insert(filename, desctinationPath, originalFilename, filetype)
}

func NewModel(repository repository.UploadImageRepository) UploadedImage {
	return uploadImpl{repository: repository}
}
