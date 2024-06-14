package token

import (
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent/repository"
	"wsw/backend/lib/utils"
)

type (
	Token interface {
		CreateToken() (*string, error)
		GetPreviewData(string) (*PreviewData, error)
	}
	PreviewData struct{}
	tokenImpl   struct {
		generator  generator.TokenGenerator
		repository repository.Token
	}
)

// GetPreviewData implements Token.
func (t tokenImpl) GetPreviewData(token string) (*PreviewData, error) {
	tokenEntity, err := t.repository.Find(token)
	if err != nil {
		return nil, err
	}
	utils.D(tokenEntity)
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
