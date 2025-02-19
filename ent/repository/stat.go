package repository

import (
	"context"
	"time"

	"wsw/backend/ent"
)

type (
	Stat interface {
		Insert(title *string, image *ent.Image) (*ent.Stat, error)
		GetImage(statEntity *ent.Stat) (*ent.Image, error)
	}

	statImp struct {
		client *ent.Client
		ctx    context.Context
	}
)

func (s *statImp) GetImage(statEntity *ent.Stat) (*ent.Image, error) {
	return statEntity.QueryImage().Only(s.ctx)
}

// Insert implements Stat.
func (s *statImp) Insert(title *string, image *ent.Image) (*ent.Stat, error) {
	statEntity, err := s.client.Stat.Create().
		SetTitle(*title).
		SetImage(image).
		SetCreatedAt(time.Now()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return statEntity, nil
}

func NewStat(client *ent.Client, ctx context.Context) Stat {
	return &statImp{client: client, ctx: ctx}
}
