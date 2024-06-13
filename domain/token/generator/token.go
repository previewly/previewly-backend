package generator

import (
	"fmt"
	"wsw/backend/lib/utils"

	"golang.org/x/exp/rand"
)

type (
	TokenGenerator interface {
		Generate() string
	}
	tokenGeneratorImpl struct {
		letters []rune
	}
)

// Generate implements TokenGenerator.
func (t *tokenGeneratorImpl) Generate() string {
	return fmt.Sprintf("%s-%s-%s", t.randomString(4), t.randomString(4), t.randomString(4))
}

func (t *tokenGeneratorImpl) randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = t.letters[rand.Intn(len(t.letters))]
	}
	return string(b)
}

func NewTokenGenerator() TokenGenerator {
	utils.InitRandom()
	return &tokenGeneratorImpl{
		letters: []rune("abcdefghijklmnopqrstuvwxyz1234567890"),
	}
}
