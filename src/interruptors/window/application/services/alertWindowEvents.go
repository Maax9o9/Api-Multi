package service

import (
    "log"
    "Multi/src/interruptors/window/domain/entities"
)

type AlertWindowUseCase interface {
    CreateWindowData(data *entities.WindowSensor) error
}

type AlertWindowService struct {
    useCase AlertWindowUseCase
}

func NewAlertWindowService(useCase AlertWindowUseCase) *AlertWindowService {
    return &AlertWindowService{
        useCase: useCase,
    }
}

func (s *AlertWindowService) CreateWindowData(data *entities.WindowSensor) error {
    err := s.useCase.CreateWindowData(data)
    if err != nil {
        log.Printf("Error al crear datos de ventana: %v", err)
        return err
    }
    return nil
}