package service

import (
    "log"
    "Multi/src/weather/domain/entities"
)

type AlertWeatherUseCase interface {
    CreateWeatherData(data *entities.SensorDataWeather) error
}

type AlertWeatherService struct {
    useCase AlertWeatherUseCase
}

func NewAlertWeatherService(useCase AlertWeatherUseCase) *AlertWeatherService {
    return &AlertWeatherService{
        useCase: useCase,
    }
}

func (s *AlertWeatherService) CreateWeatherData(data *entities.SensorDataWeather) error {
    err := s.useCase.CreateWeatherData(data)
    if err != nil {
        log.Printf("Error al crear datos meteorológicos: %v", err)
        return err
    }
    return nil
}