package repository

import (
	"context"

	"wsw/backend/ent"
)

type (
	UploadImageRepository interface {
		Insert(string, string) (*ent.UploadImage, error)
	}
	uploadrepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Inser implements UploadImageRepository.
func (u uploadrepositoryImpl) Insert(filename string, filetype string) (*ent.UploadImage, error) {
	imageEntity, err := u.client.UploadImage.Create().
		SetFilename(filename).
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
