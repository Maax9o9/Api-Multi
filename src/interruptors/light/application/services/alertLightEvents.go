package service

import (
    "log"
    "Multi/src/interruptors/light/domain/entities"
)

type AlertLightUseCase interface {
    CreateLightData(data *entities.LightData) error
}

type AlertLightService struct {
    useCase AlertLightUseCase
}

func NewAlertLightService(useCase AlertLightUseCase) *AlertLightService {
    return &AlertLightService{
        useCase: useCase,
    }
}

func (s *AlertLightService) CreateLightData(data *entities.LightData) error {
    err := s.useCase.CreateLightData(data)
    if err != nil {
        log.Printf("Error al crear datos de la luz: %v", err)
        return err
    }
    return nil
}