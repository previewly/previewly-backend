package token

import (
	"wsw/backend/domain/token/generator"
	"wsw/backend/ent/repository"
)

type (
	Token interface {
		CreateToken() (*string, error)
	}
	tokenImpl struct {
		generator  generator.TokenGenerator
		repository repository.Token
	}
)

// CreateToken implements Token.
func (t tokenImpl) CreateToken() (*string, error) {
	token, error := t.repository.InsertToken(t.generator.Generate())
	if error != nil {
		return nil, error
	}
	return &token.Value, nil
}

func NewModel(generator generator.TokenGenerator, tokenRepository repository.Token) Token {
	return tokenImpl{generator: generator, repository: tokenRepository}
}
