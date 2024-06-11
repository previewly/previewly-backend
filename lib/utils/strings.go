package utils

import (
	"fmt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func randomSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomToken() string {
	return fmt.Sprintf("%s-%s-%s", randomSeq(4), randomSeq(4), randomSeq(4))
}
