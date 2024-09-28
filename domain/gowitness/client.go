package gowitness

type (
	Client interface {
		GetUrlDetails(string) (*Details, error)
	}
	Runner interface {
		Close()
	}
	Details    struct{}
	clientImpl struct{}
	runnerImpl struct{}
)

// Close implements Runner.
func (r runnerImpl) Close() {
	panic("unimplemented")
}

func NewClient() Client {
	return clientImpl{}
}

func NewRunner() Runner {
	return runnerImpl{}
}

// GetUrlDetails implements Client.
func (c clientImpl) GetUrlDetails(string) (*Details, error) {
	panic("unimplemented")
}
