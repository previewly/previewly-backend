package token

import (
	"errors"
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent/repository"
)

type (
	Token interface {
		CreateToken() (*string, error)
		GetPreviewData(string) (*PreviewData, error)
		AddURL(string, string) (*PreviewData, error)
	}
	PreviewData struct{}
	tokenImpl   struct {
		generator  generator.TokenGenerator
		repository repository.Token
	}
)

func (t tokenImpl) isTokenExist(token string) bool {
	_, err := t.repository.Find(token)
	return err == nil
}

// AddURL implements Token.
func (t tokenImpl) AddURL(token string, url string) (*PreviewData, error) {
	if !t.isTokenExist(token) {
		return nil, errors.New("invalid token")
	}
	panic("unimplemented")
}

// GetPreviewData implements Token.
func (t tokenImpl) GetPreviewData(token string) (*PreviewData, error) {
	if !t.isTokenExist(token) {
		return nil, errors.New("invalid token")
	}
	panic("unimplemented")
}

// CreateToken implements Token.
func (t tokenImpl) CreateToken() (*string, error) {
	token, err := t.repository.InsertToken(t.generator.Generate())
	if err != nil {
		return nil, err
	}
	return &token.Value, nil
}

func NewModel(generator generator.TokenGenerator, tokenRepository repository.Token) Token {
	return tokenImpl{generator: generator, repository: tokenRepository}
}
