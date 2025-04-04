package service

import (
    "log"
    "Multi/src/hallWindow/domain/entities"
)

type AlertHallWindowUseCase interface {
    CreateHallWindow(data *entities.HallWindow) error
}

type AlertHallWindowService struct {
    useCase AlertHallWindowUseCase
}

func NewAlertHallWindowService(useCase AlertHallWindowUseCase) *AlertHallWindowService {
    return &AlertHallWindowService{
        useCase: useCase,
    }
}

func (s *AlertHallWindowService) CreateHallWindow(data *entities.HallWindow) error {
    err := s.useCase.CreateHallWindow(data)
    if err != nil {
        log.Printf("Error al crear datos de HallWindow: %v", err)
        return err
    }
    return nil
}