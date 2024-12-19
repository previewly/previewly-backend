package storage

import (
	"io"
	"os"

	"wsw/backend/domain/image/path"
)

type (
	StorageFile struct {
		NewFilename  string
		NewFilePlace string
		Directory    string
		FullPath     string
	}
	Storage interface {
		Save(string, *string, io.ReadSeeker) (*StorageFile, error)
	}
	storageImpl struct {
		nameGenerator FilenameGenerator
		pathProvider  path.PathProvider
	}
)

func NewUploadStorage(filenameGenerator FilenameGenerator, pathProvider path.PathProvider) Storage {
	return storageImpl{
		nameGenerator: filenameGenerator,
		pathProvider:  pathProvider,
	}
}

// Save implements Storage.
func (s storageImpl) Save(filename string, prefix *string, file io.ReadSeeker) (*StorageFile, error) {
	storageFile := s.createStorageFile(filename, prefix)

	// Create the uploads folder if it doesn't already exist
	err := os.MkdirAll(storageFile.Directory, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(storageFile.FullPath)
	if err != nil {
		return nil, err
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return nil, err
	}

	return storageFile, nil
}

func (s storageImpl) createStorageFile(filename string, prefix *string) *StorageFile {
	newFilename := s.nameGenerator.GenerateFilename(filename)
	newFilePlace := s.nameGenerator.GenerateFilepath(prefix)
	pathData := s.pathProvider.Provide(newFilePlace, newFilename)

	return &StorageFile{
		NewFilename:  newFilename,
		NewFilePlace: newFilePlace,
		Directory:    pathData.Directory,
		FullPath:     pathData.FullPath,
	}
}
