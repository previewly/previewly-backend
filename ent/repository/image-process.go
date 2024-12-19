package repository

import (
	"context"

	"wsw/backend/ent"
)

type (
	ImageProcessRepository     interface{}
	imageProcessRepositoryImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func NewImageProcessRepository(client *ent.Client, ctx context.Context) ImageProcessRepository {
	return imageProcessRepositoryImpl{client: client, ctx: ctx}
}
