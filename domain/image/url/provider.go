package url

type (
	Provider interface {
		Provide(filename *string) string
	}
	providerImpl struct {
		baseURL   string
		assetsURL string
	}
)

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
