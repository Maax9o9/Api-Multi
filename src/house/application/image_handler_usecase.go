package application

import (
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"Multi/src/house/domain"
)

type ImageHandlerUseCase struct {
	fileStorage domain.FileStorage
}

func NewImageHandlerUseCase(fileStorage domain.FileStorage) *ImageHandlerUseCase {
	return &ImageHandlerUseCase{
		fileStorage: fileStorage,
	}
}

func (uc *ImageHandlerUseCase) Execute(filename string, file io.Reader) (string, error) {
	filename = sanitizeFilename(filename)

	ext := filepath.Ext(filename)
	baseFilename := strings.TrimSuffix(filename, ext)
	timestamp := time.Now().UnixNano()
	newFilename := baseFilename + "_" + strconv.FormatInt(timestamp, 10) + ext

	return uc.fileStorage.SaveFile(newFilename, file)
}

func sanitizeFilename(filename string) string {
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		" ", "_",
	)
	sanitized := replacer.Replace(filename)

	if len(sanitized) > 100 {
		ext := filepath.Ext(sanitized)
		sanitized = sanitized[:100-len(ext)] + ext
	}

	return sanitized
}
