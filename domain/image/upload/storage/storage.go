package storage

import (
	"io"
	"os"
	"strings"
)

type (
	StorageFile struct {
		NewFilename  string
		NewFilePlace string
		FullPath     string
		FullFilename string
	}
	Storage interface {
		Save(string, *string, io.ReadSeeker) (*StorageFile, error)
	}
	storageImpl struct {
		destPath      string
		nameGenerator FilenameGenerator
	}
)

// Save implements Storage.
func (s storageImpl) Save(filename string, prefix *string, file io.ReadSeeker) (*StorageFile, error) {
	storageFile := s.createStorageFile(filename, prefix)

	// Create the uploads folder if it doesn't already exist
	err := os.MkdirAll(storageFile.FullPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(storageFile.FullFilename)
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
	fullPath := strings.Join([]string{
		strings.TrimSuffix(s.destPath, "/"),
		strings.TrimSuffix(newFilePlace, "/"),
	}, "/")
	fullFilename := strings.Join([]string{fullPath, newFilename}, "/")
	return &StorageFile{
		NewFilename:  newFilename,
		NewFilePlace: newFilePlace,
		FullPath:     fullPath,
		FullFilename: fullFilename,
	}
}

func NewUploadStorage(destPath string, filenameProvider FilenameGenerator) Storage {
	return storageImpl{
		destPath:      destPath,
		nameGenerator: filenameProvider,
	}
}
