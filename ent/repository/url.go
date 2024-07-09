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
		UpdateApiUrlId(*ent.Url, int) error
		Update(string, int) error
	}

	urlImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Update implements Url.
func (u *urlImpl) Update(image string, ID int) error {
	urlEntity, err := u.client.Url.Query().Where(entUrl.ID(ID)).Only(u.ctx)
	if err != nil {
		return err
	}
	_, errSave := u.client.Url.UpdateOne(urlEntity).SetImage(image).Save(u.ctx)
	return errSave
}

// UpdateApiUrlId implements Url.
func (u *urlImpl) UpdateApiUrlId(url *ent.Url, apiUrlId int) error {
	_, err := u.client.Url.UpdateOne(url).SetAPIURLID(apiUrlId).Save(u.ctx)
	return err
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
