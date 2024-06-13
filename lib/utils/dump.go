package utils

import (
	"log"

	"github.com/gookit/goutil/dump"
)

func D(vs ...any) {
	dump.P(vs)
}

func F(format string, vs ...any) {
	log.Fatalf(format, vs...)
}
