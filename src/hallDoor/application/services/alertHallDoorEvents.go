package service

import (
    "log"
    "Multi/src/hallDoor/domain/entities"
)

type AlertHallDoorUseCase interface {
    CreateHallDoor(data *entities.HallDoor) error
}

type AlertHallDoorService struct {
    useCase AlertHallDoorUseCase
}

func NewAlertHallDoorService(useCase AlertHallDoorUseCase) *AlertHallDoorService {
    return &AlertHallDoorService{
        useCase: useCase,
    }
}

func (s *AlertHallDoorService) CreateHallDoor(data *entities.HallDoor) error {
    err := s.useCase.CreateHallDoor(data)
    if err != nil {
        log.Printf("Error al crear datos de HallDoor: %v", err)
        return err
    }
    return nil
}