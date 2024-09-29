package relative

type (
	Provider interface {
		Provide(string) string
	}
	providerImpl struct{}
)

// Provide implements Provider.
func (p providerImpl) Provide(filename string) string {
	return filename
}

func NewProvider() Provider {
	return providerImpl{}
}
