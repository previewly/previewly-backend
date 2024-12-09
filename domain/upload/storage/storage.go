package storage

type (
	Storage interface {
		Save() error
	}
	storageImpl struct {
		path string
	}
)

// Save implements Storage.
func (s storageImpl) Save() error {
	return nil
}

func NewUploadStorage(path string) Storage {
	return storageImpl{
		path: path,
	}
}
