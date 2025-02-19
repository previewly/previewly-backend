package url

import "wsw/backend/domain/dto"

type (
	Provider interface {
		Provide(image dto.Image) string
	}
	providerImpl struct {
		baseURL string
	}
)

// Provide implements Provider.
func (p providerImpl) Provide(image dto.Image) string {
	return p.baseURL + image.RelativeFullPath()
}

func NewProvider(baseURL string) Provider {
	return providerImpl{baseURL: baseURL}
}
