package utils

import (
	"math/rand"
	"time"
)

func InitRandom() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
}
