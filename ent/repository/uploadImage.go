package repository

import (
	"context"

	"wsw/backend/ent"
)

type (
	UploadImageRepository interface {
		Insert(string, string, string, string) (*ent.UploadImage, error)
	}
	uploadrepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

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
