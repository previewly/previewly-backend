package repository

import (
	"context"
	"time"

	"wsw/backend/ent"
)

type (
	Stat interface {
		Insert(*string) (*ent.Stat, error)
	}

	statImp struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Insert implements Stat.
func (s *statImp) Insert(title *string) (*ent.Stat, error) {
	statEntity, err := s.client.Stat.Create().SetTitle(*title).SetCreatedAt(time.Now()).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return statEntity, nil
}

func NewStat(client *ent.Client, ctx context.Context) Stat {
	return &statImp{client: client, ctx: ctx}
}
