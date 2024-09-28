package gowitness

type (
	Client interface {
		GetUrlDetails(string)
	}
	clientImpl struct{}
)

// GetUrlDetails implements Client.
func (c clientImpl) GetUrlDetails(string) {
	panic("unimplemented")
}

func NewClient() Client {
	return clientImpl{}
}
