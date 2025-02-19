package url

import "wsw/backend/domain/dto"

type (
	Provider interface {
		Provide(fullRelativePath *string) string
		ProvideNew(image dto.Image) string
	}
	providerImpl struct {
		baseURL   string
		assetsURL string
	}
)

// ProvideNew implements Provider.
func (p providerImpl) ProvideNew(image dto.Image) string {
	return p.baseURL + image.RelativeFullPath()
}

// Provide implements Provider.
func (p providerImpl) Provide(filename *string) string {
	if filename != nil {
		return p.baseURL + *filename
	}
	return p.assetsURL + "loader-200px-200px.gif"
}

func NewProvider(baseURL string, assetsURL string) Provider {
	return providerImpl{baseURL: baseURL, assetsURL: assetsURL}
}
