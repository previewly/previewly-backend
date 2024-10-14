package repository

import (
	"context"

	"wsw/backend/ent"
)

type (
	Stat interface {
		Insert(*ent.Stat) (*ent.Stat, error)
	}

	statImp struct {
		client *ent.Client
		ctx    context.Context
	}
)

// Insert implements Stat.
func (s *statImp) Insert(*ent.Stat) (*ent.Stat, error) {
	panic("unimplemented")
}

func NewStat(client *ent.Client, ctx context.Context) Stat {
	return &statImp{client: client, ctx: ctx}
}
