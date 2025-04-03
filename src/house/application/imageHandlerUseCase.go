package application

import (
    "io"
    "log"
    "os"
    "path/filepath"
)

type ImageHandlerUseCase struct{}

func NewImageHandlerUseCase() *ImageHandlerUseCase {
    return &ImageHandlerUseCase{}
}

func (uc *ImageHandlerUseCase) SaveImage(imageName string, imageFile io.Reader) (string, error) {
    imageDir := "Multi/images"
    if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
        log.Printf("Error al crear el directorio de imágenes: %v", err)
        return "", err
    }

    imagePath := filepath.Join(imageDir, imageName)
    outFile, err := os.Create(imagePath)
    if err != nil {
        log.Printf("Error al crear el archivo de imagen: %v", err)
        return "", err
    }
    defer outFile.Close()

    if _, err := io.Copy(outFile, imageFile); err != nil {
        log.Printf("Error al guardar la imagen: %v", err)
        return "", err
    }

    log.Printf("Imagen guardada exitosamente en: %s", imagePath)
    return imagePath, nil
}