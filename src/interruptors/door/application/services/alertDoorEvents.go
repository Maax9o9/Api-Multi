package service

import (
    "log"
    "Multi/src/interruptors/door/domain/entities"
)

type AlertDoorUseCase interface {
    Create(data *entities.DoorData) error
}

type AlertDoorService struct {
    useCase AlertDoorUseCase
}

func NewAlertDoorService(useCase AlertDoorUseCase) *AlertDoorService {
    return &AlertDoorService{
        useCase: useCase,
    }
}

func (s *AlertDoorService) CreateDoorData(data *entities.DoorData) error {
    err := s.useCase.Create(data)
    if err != nil {
        log.Printf("Error al crear datos de la puerta: %v", err)
        return err
    }
    return nil
}