package repository

import (
	"context"

	"wsw/backend/ent"
	"wsw/backend/ent/types"
	entImage "wsw/backend/ent/uploadimage"
)

type (
	UploadImageRepository interface {
		Insert(filename string, destinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.UploadImage, error)
		CreateProcess(entity *ent.UploadImage, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error)
		GetByID(int) (*ent.UploadImage, error)
	}
	uploadrepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func (u uploadrepositoryImpl) CreateProcess(imageEntity *ent.UploadImage, imageProcesses []types.ImageProcess, hash string) (*ent.ImageProcess, error) {
	process, err := u.client.ImageProcess.Create().
		SetStatus(types.Pending).
		SetProcessHash(hash).
		SetProcesses(imageProcesses).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	_, errUpdated := u.client.UploadImage.
		UpdateOne(imageEntity).
		AddImageprocess(process).
		Save(u.ctx)
	if errUpdated != nil {
		return nil, errUpdated
	}
	return process, nil
}

// GetByID implements UploadImageRepository.
func (u uploadrepositoryImpl) GetByID(imageID int) (*ent.UploadImage, error) {
	return u.client.UploadImage.Query().Where(entImage.ID(imageID)).Only(u.ctx)
}

// Inser implements UploadImageRepository.
func (u uploadrepositoryImpl) Insert(filename string, destinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.UploadImage, error) {
	entity := u.client.UploadImage.Create().
		SetFilename(filename).
		SetOriginalFilename(originalFilename).
		SetDestinationPath(destinationPath).
		SetType(filetype)

	if extraValue != nil {
		entity = entity.SetExtraValue(*extraValue)
	}
	imageEntity, err := entity.Save(u.ctx)
	if err != nil {
		return nil, err
	}
	return imageEntity, nil
}

func NewUploadImageRepository(client *ent.Client, ctx context.Context) UploadImageRepository {
	return uploadrepositoryImpl{client: client, ctx: ctx}
}
