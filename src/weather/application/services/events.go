package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/weather/domain/entities"
)

type WeatherUseCase interface {
    GetAllWeatherData() ([]entities.SensorDataWeather, error)
    GetWeatherDataByID(id int) (*entities.SensorDataWeather, error)
}

type WeatherService struct {
    useCase     WeatherUseCase
    latestData  *entities.SensorDataWeather
}

func NewWeatherService(useCase WeatherUseCase) *WeatherService {
    return &WeatherService{
        useCase: useCase,
    }
}

func (s *WeatherService) GetLatestWeatherData() (*entities.SensorDataWeather, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("No hay datos recientes")
    }
    return s.latestData, nil
}

func (s *WeatherService) GetAllWeatherData() ([]entities.SensorDataWeather, error) {
    data, err := s.useCase.GetAllWeatherData()
    if err != nil {
        log.Printf("Error al obtener todos los datos meteorológicos: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *WeatherService) UpdateLatestWeatherData() error {
    data, err := s.useCase.GetAllWeatherData()
    if err != nil {
        log.Printf("Error al actualizar los datos meteorológicos: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1] 
        log.Printf("Últimos datos meteorológicos actualizados: %+v", s.latestData)
    }
    return nil
}

func (s *WeatherService) SerializeWeatherData(data []entities.SensorDataWeather) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos meteorológicos: %v", err)
        return "", err
    }
    return string(jsonData), nil
}