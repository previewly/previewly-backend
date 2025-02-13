package image

import (
	"io"

	"wsw/backend/domain/image/upload/storage"
	"wsw/backend/ent"
	"wsw/backend/model/image"

	"github.com/xorcare/pointer"
)

type (
	Saver interface {
		SaveImage(imageName string, image io.ReadSeeker, contentType string, extraValue *string) (*ent.UploadImage, error)
	}
	saverImpl struct {
		storage storage.Storage
		model   image.UploadedImage
	}
)

func NewSaver(model image.UploadedImage, storage storage.Storage) Saver {
	return saverImpl{storage: storage, model: model}
}

func (s saverImpl) SaveImage(imageName string, image io.ReadSeeker, contentType string, extraValue *string) (*ent.UploadImage, error) {
	storageFile, err := s.storage.Save(imageName, pointer.String("o/"), image)
	if err != nil {
		return nil, err
	}
	return s.model.Insert(storageFile.NewFilename, storageFile.NewFilePlace, imageName, contentType, extraValue)
}
