package application

import (
    "Multi/src/house/domain/entities"
    "io"
    "log"
)

type AddHouseUseCase struct {
    imageHandler    *ImageHandlerUseCase
    locationHandler *LocationHandlerUseCase
}

func NewAddHouseUseCase(imageHandler *ImageHandlerUseCase, locationHandler *LocationHandlerUseCase) *AddHouseUseCase {
    return &AddHouseUseCase{
        imageHandler:    imageHandler,
        locationHandler: locationHandler,
    }
}

func (uc *AddHouseUseCase) AddHouse(house *entities.HouseProfile, imageName string, imageFile io.Reader) error {
    imagePath, err := uc.imageHandler.SaveImage(imageName, imageFile)
    if err != nil {
        return err
    }
    house.Image = imagePath

    if err := uc.locationHandler.SaveHouseLocation(house); err != nil {
        return err
    }

    log.Printf("Casa creada exitosamente con imagen y ubicación: %+v", house)
    return nil
}