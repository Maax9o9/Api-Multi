package application

import (
    "Multi/src/weather/domain"
    "Multi/src/weather/domain/entities"
    "Multi/src/weather/application/repositorys"
    "encoding/json"
    "log"
)

type AlertWeatherUseCase struct {
    repo       domain.WeatherRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertWeatherUseCase(repo domain.WeatherRepository, rabbitRepo *repositorys.RabbitRepository) *AlertWeatherUseCase {
    return &AlertWeatherUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertWeatherUseCase) CreateWeatherData(data *entities.SensorDataWeather) error {
    return uc.repo.Create(data)
}

func (uc *AlertWeatherUseCase) ProcessWeatherData(message []byte) error {
    var weatherData entities.SensorDataWeather
    err := json.Unmarshal(message, &weatherData)
    if err != nil {
        log.Printf("Error unmarshalling weather data: %v", err)
        return err
    }

    err = uc.CreateWeatherData(&weatherData)
    if err != nil {
        log.Printf("Error saving weather data: %v", err)
        return err
    }

    log.Printf("Weather data processed and saved: %+v", weatherData)
    return nil
}