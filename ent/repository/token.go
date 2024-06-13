package repository

type Token interface{}

type tokenImpl struct{}

func NewToken() Token {
	return &tokenImpl{}
}
