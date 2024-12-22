package url

import (
	"wsw/backend/app/config"
	"wsw/backend/domain/url/screenshot"
)

type (
	Provider interface {
		Provide(filename *string) string
	}
	providerImpl struct {
		baseURL string
		loader  screenshot.Loader
	}
)

// Provide implements Provider.
func (p providerImpl) Provide(filename *string) string {
	if filename != nil {
		return p.baseURL + *filename
	}
	return p.loader.Provide()
}

func NewProvider(config config.Gowitness, loader screenshot.Loader) Provider {
	return providerImpl{
		baseURL: config.ScreenshotBaseUrl,
		loader:  loader,
	}
}
