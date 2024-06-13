package token

import "wsw/backend/lib/utils"

type (
	Token     interface{ CreateToken() (string, error) }
	tokenImpl struct{}
)

// CreateToken implements Token.
func (t tokenImpl) CreateToken() (string, error) {
	// todo
	utils.InitRandom()
	return utils.RandomToken(), nil
}

func NewModel() Token {
	return tokenImpl{}
}
