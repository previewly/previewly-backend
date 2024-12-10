package upload

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
)

type (
	UploadImage interface {
		Insert(string, string) (*ent.UploadImage, error)
	}
	uploadImpl struct {
		repository repository.UploadImageRepository
	}
)

// Insert implements UploadImage.
func (u uploadImpl) Insert(filename string, filetype string) (*ent.UploadImage, error) {
	return u.repository.Insert(filename, filetype)
}

func NewModel(repository repository.UploadImageRepository) UploadImage {
	return uploadImpl{repository: repository}
}
