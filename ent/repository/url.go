package repository

import (
	"context"

	"wsw/backend/domain/url"
	"wsw/backend/ent"
	entUrl "wsw/backend/ent/url"
)

type (
	Url interface {
		TryGet(string) *ent.Url
		Insert(string) (*ent.Url, error)
		Update(string, url.Status, int, error) (*ent.Url, error)
	}

	urlImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Update implements Url.
func (u *urlImpl) Update(relativePath string, status url.Status, ID int, _ error) (*ent.Url, error) {
	urlEntity, err := u.client.Url.Query().Where(entUrl.ID(ID)).Only(u.ctx)
	if err != nil {
		return nil, err
	}
	return u.client.Url.UpdateOne(urlEntity).SetStatus(status).SetRelativePath(relativePath).Save(u.ctx)
}

// Insert implements Url.
func (u *urlImpl) Insert(url string) (*ent.Url, error) {
	return u.client.Url.Create().SetURL(url).SetStatus("pending").Save(u.ctx)
}

// TryGet implements Url.
func (u *urlImpl) TryGet(url string) *ent.Url {
	entity, _ := u.client.Url.Query().Where(entUrl.URL(url)).Only(u.ctx)
	return entity
}

func NewUrl(client *ent.Client, ctx context.Context) Url {
	return &urlImpl{client: client, ctx: ctx}
}
