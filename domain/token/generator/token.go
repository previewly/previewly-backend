package generator

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type (
	TokenGenerator interface {
		Generate() string
	}
	tokenGeneratorImpl struct{}
)

// Generate implements TokenGenerator.
func (t *tokenGeneratorImpl) Generate() string {
	splitted := strings.Split(uuid.New().String(), "-")
	return fmt.Sprintf("%s-%s-%s", splitted[1], splitted[2], splitted[3])
}

func NewTokenGenerator() TokenGenerator {
	return &tokenGeneratorImpl{}
}
