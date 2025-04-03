package service

import (
    "log"
    "Multi/src/interruptors/gas/domain/entities"
)

type AlertGasUseCase interface {
    CreateGasData(data *entities.GasSensor) error
}

type AlertGasService struct {
    useCase AlertGasUseCase
}

func NewAlertGasService(useCase AlertGasUseCase) *AlertGasService {
    return &AlertGasService{
        useCase: useCase,
    }
}

func (s *AlertGasService) CreateGasData(data *entities.GasSensor) error {
    err := s.useCase.CreateGasData(data)
    if err != nil {
        log.Printf("Error al crear datos de gas: %v", err)
        return err
    }
    return nil
}