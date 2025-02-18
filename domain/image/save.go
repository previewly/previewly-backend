package image

import (
	"io"

	"wsw/backend/domain/image/storage"
	"wsw/backend/ent"
	"wsw/backend/model/image"

	"github.com/xorcare/pointer"
)

type (
	Saver interface {
		SaveImage(imageName string, image io.ReadSeeker, contentType string, extraValue *string) (*ent.Image, error)
	}
	saverImpl struct {
		storage storage.Storage
		model   image.Model
	}
)

func NewSaver(model image.Model, storage storage.Storage) Saver {
	return saverImpl{storage: storage, model: model}
}

func (s saverImpl) SaveImage(imageName string, image io.ReadSeeker, contentType string, extraValue *string) (*ent.Image, error) {
	storageFile, err := s.storage.Save(imageName, pointer.String("o/"), image)
	if err != nil {
		return nil, err
	}
	return s.model.Insert(storageFile.NewFilename, storageFile.NewFilePlace, imageName, contentType, extraValue)
}
