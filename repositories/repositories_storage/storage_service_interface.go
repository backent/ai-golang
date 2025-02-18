package repositories_storage

import "mime/multipart"

type StorageServiceInterface interface {
	SaveFile(file multipart.File, fileName string, destination string) error
}
