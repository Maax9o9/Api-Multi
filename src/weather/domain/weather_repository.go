package domain

import "Multi/src/weather/domain/entities"

type WeatherRepository interface {
    Create(data *entities.SensorDataWeather) error
    GetAll() ([]entities.SensorDataWeather, error)
    GetByID(id int) (*entities.SensorDataWeather, error)
}