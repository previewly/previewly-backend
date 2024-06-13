package token

import (
	"wsw/backend/domain/token/generator"
)

type (
	Token     interface{ CreateToken() (string, error) }
	tokenImpl struct {
		generator generator.TokenGenerator
	}
)

// CreateToken implements Token.
func (t tokenImpl) CreateToken() (string, error) {
	return t.generator.Generate(), nil
}

func NewModel(generator generator.TokenGenerator) Token {
	return tokenImpl{generator: generator}
}
