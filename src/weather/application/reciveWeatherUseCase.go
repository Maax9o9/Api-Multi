package application

import (
    "Multi/src/weather/domain"
    "Multi/src/weather/domain/entities"
)

type ReceiveWeatherUseCase struct {
    repo domain.WeatherRepository
}

func NewReceiveWeatherUseCase(repo domain.WeatherRepository) *ReceiveWeatherUseCase {
    return &ReceiveWeatherUseCase{
        repo: repo,
    }
}

func (uc *ReceiveWeatherUseCase) GetAllWeatherData() ([]entities.SensorDataWeather, error) {
    return uc.repo.GetAll()
}

func (uc *ReceiveWeatherUseCase) GetWeatherDataByID(id int) (*entities.SensorDataWeather, error) {
    return uc.repo.GetByID(id)
}