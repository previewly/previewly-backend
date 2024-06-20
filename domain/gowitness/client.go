package gowitness

type (
	Client interface {
		AddUrl(string)
		Search(string)
	}
	clientImpl struct{}
)

// AddUrl implements Client.
func (c *clientImpl) AddUrl(string) {
	panic("gowitness.AddUrl is unimplemented")
}

// Search implements Client.
func (c *clientImpl) Search(string) {
	panic("gowitness.AddUrl is unimplemented")
}

func NewClient() Client {
	return &clientImpl{}
}
