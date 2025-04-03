package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
    "io"
    "log"
    "os"
    "path/filepath"
)

type EditHouseUseCase struct {
    repo domain.HouseRepository
}

func NewEditHouseUseCase(repo domain.HouseRepository) *EditHouseUseCase {
    return &EditHouseUseCase{
        repo: repo,
    }
}

func (uc *EditHouseUseCase) UpdateHouse(house *entities.HouseProfile, imageName string, imageFile io.Reader) error {
    if imageName != "" && imageFile != nil {
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
    }

    if err := uc.repo.Update(house); err != nil {
        log.Printf("Error al actualizar la casa en el repositorio: %v", err)
        return err
    }

    log.Printf("Casa actualizada exitosamente: %+v", house)
    return nil
}