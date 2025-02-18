package repository

import (
	"context"

	"wsw/backend/ent"
	entImage "wsw/backend/ent/image"
	entImageProcess "wsw/backend/ent/imageprocess"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"
)

type (
	ImageProcessRepository interface {
		Update(*ent.ImageProcess, string, types.StatusEnum, *string) (*ent.ImageProcess, error)
		GetByHash(imageID int, processesHash string) (*ent.ImageProcess, error)
	}
	imageProcessRepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func (i imageProcessRepositoryImpl) GetByHash(imageID int, processesHash string) (*ent.ImageProcess, error) {
	return i.client.ImageProcess.Query().
		Where(
			entImageProcess.HasUploadimageWith(entImage.ID(imageID)),
			entImageProcess.ProcessHash(processesHash),
		).
		Only(i.ctx)
}

func (i imageProcessRepositoryImpl) Update(processEntity *ent.ImageProcess, prefix string, status types.StatusEnum, err *string) (*ent.ImageProcess, error) {
	return i.client.ImageProcess.UpdateOne(processEntity).
		SetPathPrefix(prefix).
		SetStatus(status).
		SetError(utils.ToString(err)).
		Save(i.ctx)
}

func NewImageProcessRepository(client *ent.Client, ctx context.Context) ImageProcessRepository {
	return imageProcessRepositoryImpl{client: client, ctx: ctx}
}
