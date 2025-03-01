package di

import (
	"wsw/backend/lib/utils"

	"github.com/golobby/container/v3"
)

func InitModule(name string, resolvers ...interface{}) {
	for _, resolver := range resolvers {
		err := container.Singleton(resolver)
		if err != nil {
			utils.F("Couldnt inititalize service: %v of module %s", err, name)
		}
	}
}
