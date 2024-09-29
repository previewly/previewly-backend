package screenshot

import "wsw/backend/app/config"

type (
	Provider interface {
		Provide(filename *string) string
	}
	providerImpl struct {
		baseURL string
	}
)

// Provide implements Provider.
func (p providerImpl) Provide(filename *string) string {
	if filename != nil {
		return p.baseURL + *filename
	}
	// FIXME
	return "loader"
}

func NewProvider(config config.Gowitness) Provider {
	return providerImpl{
		baseURL: config.ScreenshotBaseUrl,
	}
}
