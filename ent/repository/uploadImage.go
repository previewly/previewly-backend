package repository

import (
	"context"

	"wsw/backend/ent"
	entImage "wsw/backend/ent/uploadimage"
)

type (
	UploadImageRepository interface {
		Insert(string, string, string, string) (*ent.UploadImage, error)
		GetByID(int) (*ent.UploadImage, error)
	}
	uploadrepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// GetByID implements UploadImageRepository.
func (u uploadrepositoryImpl) GetByID(imageID int) (*ent.UploadImage, error) {
	return u.client.UploadImage.Query().Where(entImage.ID(imageID)).Only(u.ctx)
}

// Inser implements UploadImageRepository.
func (u uploadrepositoryImpl) Insert(filename string, destinationPath string, originalFilename string, filetype string) (*ent.UploadImage, error) {
	imageEntity, err := u.client.UploadImage.Create().
		SetFilename(filename).
		SetOriginalFilename(originalFilename).
		SetDestinationPath(destinationPath).
		SetType(filetype).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}
	return imageEntity, nil
}

func NewUploadImageRepository(client *ent.Client, ctx context.Context) UploadImageRepository {
	return uploadrepositoryImpl{client: client, ctx: ctx}
}
