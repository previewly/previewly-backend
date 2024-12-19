package repository

import (
	"context"

	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"
)

type (
	ImageProcessRepository interface {
		Update(*ent.ImageProcess, *string, types.StatusEnum, *string) (*ent.ImageProcess, error)
	}
	imageProcessRepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func (i imageProcessRepositoryImpl) Update(processEntity *ent.ImageProcess, prefix *string, status types.StatusEnum, err *string) (*ent.ImageProcess, error) {
	return i.client.ImageProcess.UpdateOne(processEntity).
		SetPathPrefix(*prefix).
		SetStatus(status).
		SetError(utils.ToString(err)).
		Save(i.ctx)
}

func NewImageProcessRepository(client *ent.Client, ctx context.Context) ImageProcessRepository {
	return imageProcessRepositoryImpl{client: client, ctx: ctx}
}
