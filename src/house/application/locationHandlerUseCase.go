package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
    "log"
)

type LocationHandlerUseCase struct {
    repo domain.HouseRepository
}

func NewLocationHandlerUseCase(repo domain.HouseRepository) *LocationHandlerUseCase {
    return &LocationHandlerUseCase{
        repo: repo,
    }
}

func (uc *LocationHandlerUseCase) SaveHouseLocation(house *entities.HouseProfile) error {
    if err := uc.repo.Create(house); err != nil {
        log.Printf("Error al guardar la casa en el repositorio: %v", err)
        return err
    }

    log.Printf("Ubicación de la casa guardada exitosamente: %+v", house)
    return nil
}