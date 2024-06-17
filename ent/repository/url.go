package repository

import (
	"context"
	"wsw/backend/ent"
)

type (
	Url interface {
		TryGet(string) *ent.Url
		Insert(string) (*ent.Url, error)
	}

	urlImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Insert implements Url.
func (u *urlImpl) Insert(string) (*ent.Url, error) {
	panic("unimplemented")
}

// TryGet implements Url.
func (u *urlImpl) TryGet(string) *ent.Url {
	panic("unimplemented")
}

func NewUrl(client *ent.Client, ctx context.Context) Url {
	return &urlImpl{client: client, ctx: ctx}
}
