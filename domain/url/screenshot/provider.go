package screenshot

import (
	"wsw/backend/app/config"
)

type (
	Provider interface {
		Provide(filename *string) string
	}
	providerImpl struct {
		baseURL string
		loader  Loader
	}
)

// Provide implements Provider.
func (p providerImpl) Provide(filename *string) string {
	if filename != nil {
		return p.baseURL + *filename
	}
	return p.loader.Provide()
}

func NewProvider(config config.Gowitness, loader Loader) Provider {
	return providerImpl{
		baseURL: config.ScreenshotBaseUrl,
		loader:  loader,
	}
}
