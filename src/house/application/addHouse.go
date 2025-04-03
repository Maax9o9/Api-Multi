package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
    "io"
    "log"
    "os"
    "path/filepath"
)

type AddHouseUseCase struct {
    repo domain.HouseRepository
}

func NewAddHouseUseCase(repo domain.HouseRepository) *AddHouseUseCase {
    return &AddHouseUseCase{
        repo: repo,
    }
}

func (uc *AddHouseUseCase) AddHouse(house *entities.HouseProfile, imageName string, imageFile io.Reader) error {
    imageDir := "Multi/images"
    if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
        log.Printf("Error al crear el directorio de imágenes: %v", err)
        return err
    }

    imagePath := filepath.Join(imageDir, imageName)
    outFile, err := os.Create(imagePath)
    if err != nil {
        log.Printf("Error al crear el archivo de imagen: %v", err)
        return err
    }
    defer outFile.Close()

    if _, err := io.Copy(outFile, imageFile); err != nil {
        log.Printf("Error al guardar la imagen: %v", err)
        return err
    }

    house.Image = imagePath

    if err := uc.repo.Create(house); err != nil {
        log.Printf("Error al guardar la casa en el repositorio: %v", err)
        return err
    }

    log.Printf("Casa creada exitosamente con imagen: %+v", house)
    return nil
}