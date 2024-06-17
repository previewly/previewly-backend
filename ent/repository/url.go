package repository

import (
	"context"
	"wsw/backend/ent"
	entUrl "wsw/backend/ent/url"
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
func (u *urlImpl) Insert(url string) (*ent.Url, error) {
	return u.client.Url.Create().SetURL(url).Save(u.ctx)
}

// TryGet implements Url.
func (u *urlImpl) TryGet(url string) *ent.Url {
	entity, _ := u.client.Url.Query().Where(entUrl.URL(url)).Only(u.ctx)
	return entity
}

func NewUrl(client *ent.Client, ctx context.Context) Url {
	return &urlImpl{client: client, ctx: ctx}
}
