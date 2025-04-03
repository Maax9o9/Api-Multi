package service

import (
    "log"
    "Multi/src/interruptors/movement/domain/entities"
)

type AlertMovementUseCase interface {
    Create(data *entities.MotionSensor) error
}

type AlertMovementService struct {
    useCase AlertMovementUseCase
}

func NewAlertMovementService(useCase AlertMovementUseCase) *AlertMovementService {
    return &AlertMovementService{
        useCase: useCase,
    }
}

func (s *AlertMovementService) CreateMovementData(data *entities.MotionSensor) error {
    err := s.useCase.Create(data)
    if err != nil {
        log.Printf("Error al crear datos de movimiento: %v", err)
        return err
    }
    return nil
}