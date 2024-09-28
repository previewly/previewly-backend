package gowitness

type (
	Client interface {
		GetUrlDetails(string) (*Details, error)
	}
	Details    struct{}
	clientImpl struct{}
)

// GetUrlDetails implements Client.
func (c clientImpl) GetUrlDetails(string) (*Details, error) {
	panic("unimplemented")
}

func NewClient() Client {
	return clientImpl{}
}
