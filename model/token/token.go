package token

import (
	"wsw/backend/domain/preview"
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent/repository"
)

type (
	Token interface {
		CreateToken() (*string, error)
		GetPreviewData(string) (*preview.PreviewData, error)
		IsTokenExist(string) bool
	}
	tokenImpl struct {
		generator  generator.TokenGenerator
		repository repository.Token
	}
)

func (t tokenImpl) IsTokenExist(token string) bool {
	_, err := t.repository.Find(token)
	return err == nil
}

// GetPreviewData implements Token.
func (t tokenImpl) GetPreviewData(token string) (*preview.PreviewData, error) {
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
