package application

import (
    "Multi/src/weather/domain"
    "Multi/src/weather/domain/entities"
)

type ReceiveWeatherByIDUseCase struct {
    repo domain.WeatherRepository
}

func NewReceiveWeatherByIDUseCase(repo domain.WeatherRepository) *ReceiveWeatherByIDUseCase {
    return &ReceiveWeatherByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveWeatherByIDUseCase) GetWeatherDataByID(id int) (*entities.SensorDataWeather, error) {
    return uc.repo.GetByID(id)
}