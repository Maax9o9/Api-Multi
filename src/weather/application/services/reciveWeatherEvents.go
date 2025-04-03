package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/weather/domain/entities"
)

type GetAllWeatherUseCase interface {
    GetAllWeatherData() ([]entities.SensorDataWeather, error)
}

type GetWeatherByIDUseCase interface {
    GetWeatherDataByID(id int) (*entities.SensorDataWeather, error)
}

type ReceiveWeatherService struct {
    getAllUseCase  GetAllWeatherUseCase
    getByIDUseCase GetWeatherByIDUseCase
    latestData     *entities.SensorDataWeather
}

func NewReceiveWeatherService(getAllUseCase GetAllWeatherUseCase, getByIDUseCase GetWeatherByIDUseCase) *ReceiveWeatherService {
    return &ReceiveWeatherService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveWeatherService) GetAllWeatherData() ([]entities.SensorDataWeather, error) {
    data, err := s.getAllUseCase.GetAllWeatherData()
    if err != nil {
        log.Printf("Error al obtener todos los datos meteorológicos: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveWeatherService) GetWeatherDataByID(id int) (*entities.SensorDataWeather, error) {
    data, err := s.getByIDUseCase.GetWeatherDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos meteorológicos por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveWeatherService) GetLatestWeatherData() (*entities.SensorDataWeather, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes")
    }
    return s.latestData, nil
}

func (s *ReceiveWeatherService) UpdateLatestWeatherData() error {
    data, err := s.getAllUseCase.GetAllWeatherData()
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

func (s *ReceiveWeatherService) SerializeWeatherData(data []entities.SensorDataWeather) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos meteorológicos: %v", err)
        return "", err
    }
    return string(jsonData), nil
}