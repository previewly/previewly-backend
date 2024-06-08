package utils

import (
	"log"

	"github.com/gookit/goutil/dump"
)

func D(vs ...any) {
	dump.P(vs)
}

func F(vs ...any) {
	log.Fatal(vs...)
}
