package repositories_storage

import (
	"io"
	"mime/multipart"
	"os"
)

type StorageServiceLocalImplementation struct {
}

func NewStorageServiceLocalImplementation() StorageServiceInterface {
	return &StorageServiceLocalImplementation{}
}

func (implementation *StorageServiceLocalImplementation) SaveFile(file multipart.File, fileName string, destination string) error {
	dst, err := os.Create(destination + "/" + fileName)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy the file data to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
