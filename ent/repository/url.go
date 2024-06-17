package repository

import (
	"context"
	"wsw/backend/ent"
)

type (
	Url     interface{}
	urlImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

func NewUrl(client *ent.Client, ctx context.Context) Url {
	return &urlImpl{client: client, ctx: ctx}
}
