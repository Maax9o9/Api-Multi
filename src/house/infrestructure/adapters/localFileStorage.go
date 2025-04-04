package adapters

import (
	"io"
	"os"
	"path/filepath"

	"Multi/src/house/domain"
)

type LocalFileStorage struct {
	basePath string
}

func NewLocalFileStorage(basePath string) domain.FileStorage {
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		os.MkdirAll(basePath, 0755)
	}

	return &LocalFileStorage{
		basePath: basePath,
	}
}

func (fs *LocalFileStorage) SaveFile(filename string, file io.Reader) (string, error) {
	fullPath := filepath.Join(fs.basePath, filename)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return "/uploads/" + filename, nil
}

func (fs *LocalFileStorage) DeleteFile(filePath string) error {
	filename := filepath.Base(filePath)
	fullPath := filepath.Join(fs.basePath, filename)
	return os.Remove(fullPath)
}
