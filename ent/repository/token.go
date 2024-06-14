package repository

import (
	"context"
	"wsw/backend/ent"
	entToken "wsw/backend/ent/token"
)

type Token interface {
	InsertToken(string) (*ent.Token, error)
	Find(token string) (*ent.Token, error)
}

type tokenImpl struct {
	ctx    context.Context
	client *ent.Client
}

func (t *tokenImpl) Find(token string) (*ent.Token, error) {
	return t.client.Token.Query().Where(entToken.Value(token)).Only(t.ctx)
}

// InsertToken implements Token.
func (t *tokenImpl) InsertToken(token string) (*ent.Token, error) {
	tokenEntity, err := t.client.Token.Create().SetValue(token).Save(t.ctx)
	if err != nil {
		return nil, err
	}
	return tokenEntity, nil
}

func NewToken(client *ent.Client, ctx context.Context) Token {
	return &tokenImpl{client: client, ctx: ctx}
}
