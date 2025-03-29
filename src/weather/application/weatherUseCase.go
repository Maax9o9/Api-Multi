package application

import (
    "Multi/src/weather/domain/entities"
    "Multi/src/weather/domain"
    "Multi/src/weather/application/repositorys"
)

type WeatherUseCase struct {
    repo        domain.WeatherRepository
    rabbitRepo  *repositorys.RabbitRepository
}

func NewWeatherUseCase(repo domain.WeatherRepository, rabbitRepo *repositorys.RabbitRepository) *WeatherUseCase {
    return &WeatherUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *WeatherUseCase) GetAllWeatherData() ([]entities.SensorDataWeather, error) {
    return uc.repo.GetAll()
}

func (uc *WeatherUseCase) GetWeatherDataByID(id int) (*entities.SensorDataWeather, error) {
    return uc.repo.GetByID(id)
}

func (uc *WeatherUseCase) ProcessWeatherData(processMessage func(body []byte)) error {
    return uc.rabbitRepo.ProcessWeatherData(processMessage)
}