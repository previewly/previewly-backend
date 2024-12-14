package repository

import (
	"context"

	"wsw/backend/domain/url"
	"wsw/backend/ent"
	entUrl "wsw/backend/ent/url"
)

type (
	Url interface {
		Get(string) (*ent.Url, error)
		TryGet(string) *ent.Url
		Insert(string) (*ent.Url, error)

		SaveFailure(string, int) (*ent.Url, error)
		SaveSuccess(string, *ent.Stat, int) (*ent.Url, error)

		GetErrors(*ent.Url) ([]*ent.ErrorResult, error)
		GetStats(*ent.Url) ([]*ent.Stat, error)
	}

	urlImpl struct {
		client *ent.Client
		ctx    context.Context
	}
)

// SaveFailure implements Url.
func (u *urlImpl) SaveFailure(errorMessage string, ID int) (*ent.Url, error) {
	urlEntity, err := u.client.Url.Query().Where(entUrl.ID(ID)).Only(u.ctx)
	if err != nil {
		return nil, err
	}
	return u.client.Url.UpdateOne(urlEntity).
		SetStatus(url.Error).
		Save(u.ctx)
}

// SaveSuccess implements Url.
func (u *urlImpl) SaveSuccess(relativePath string, statEntity *ent.Stat, ID int) (*ent.Url, error) {
	urlEntity, err := u.client.Url.Query().Where(entUrl.ID(ID)).Only(u.ctx)
	if err != nil {
		return nil, err
	}
	return u.client.Url.UpdateOne(urlEntity).
		SetStatus(url.Success).
		SetRelativePath(relativePath).
		AddStat(statEntity).
		Save(u.ctx)
}

func (u *urlImpl) GetErrors(entity *ent.Url) ([]*ent.ErrorResult, error) {
	return entity.QueryErrorresult().All(u.ctx)
}

func (u *urlImpl) GetStats(entity *ent.Url) ([]*ent.Stat, error) {
	return entity.QueryStat().All(u.ctx)
}

// Get implements Url.
func (u *urlImpl) Get(url string) (*ent.Url, error) {
	return u.client.Url.Query().Where(entUrl.URL(url)).Only(u.ctx)
}

// Update implements Url.
func (u *urlImpl) Update(relativePath string, status url.Status, ID int) (*ent.Url, error) {
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
	entity, _ := u.Get(url)
	return entity
}

func NewUrl(client *ent.Client, ctx context.Context) Url {
	return &urlImpl{client: client, ctx: ctx}
}
