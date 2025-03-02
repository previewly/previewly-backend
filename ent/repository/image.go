package repository

import (
	"context"

	"wsw/backend/ent"
	entImage "wsw/backend/ent/image"
	"wsw/backend/ent/types"
)

type (
	ImageRepository interface {
		Insert(filename string, destinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error)
		CreateProcess(entity *ent.Image, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error)
		GetByID(int) (*ent.Image, error)
	}
	imageRepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func (u imageRepositoryImpl) CreateProcess(imageEntity *ent.Image, imageProcesses []types.ImageProcess, hash string) (*ent.ImageProcess, error) {
	process, err := u.client.ImageProcess.Create().
		SetStatus(types.Pending).
		SetProcessHash(hash).
		SetProcesses(imageProcesses).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	_, errUpdated := u.client.Image.
		UpdateOne(imageEntity).
		AddImageprocess(process).
		Save(u.ctx)
	if errUpdated != nil {
		process.Update().SetError(errUpdated.Error()).Save(u.ctx)
		return nil, errUpdated
	}
	return process, nil
}

// GetByID implements UploadImageRepository.
func (u imageRepositoryImpl) GetByID(imageID int) (*ent.Image, error) {
	return u.client.Image.Query().Where(entImage.ID(imageID)).Only(u.ctx)
}

// Inser implements UploadImageRepository.
func (u imageRepositoryImpl) Insert(filename string, destinationPath string, originalFilename string, filetype string, extraValue *string) (*ent.Image, error) {
	entity := u.client.Image.Create().
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

func NewUploadImageRepository(client *ent.Client, ctx context.Context) ImageRepository {
	return imageRepositoryImpl{client: client, ctx: ctx}
}
