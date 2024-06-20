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
	panic("unimplemented")
}

// Search implements Client.
func (c *clientImpl) Search(string) {
	panic("unimplemented")
}

func NewClient() Client {
	return &clientImpl{}
}
