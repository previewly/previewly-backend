package screenshot

type (
	Loader interface {
		Provide() string
	}
	loaderImpl struct {
		baseURL string
	}
)

// Provide implements Loader.
func (l loaderImpl) Provide() string {
	return l.baseURL + "loader-200px-200px.gif"
}

func NewLoader(baseURL string) Loader {
	return loaderImpl{baseURL: baseURL}
}
