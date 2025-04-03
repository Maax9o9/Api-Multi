package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/window/domain/entities"
)

type ReceiveWindowUseCase interface {
    GetAllWindowData() ([]entities.WindowSensor, error)
    GetWindowDataByID(id int) (*entities.WindowSensor, error)
}

type ReceiveWindowService struct {
    useCase    ReceiveWindowUseCase
    latestData *entities.WindowSensor
}

func NewReceiveWindowService(useCase ReceiveWindowUseCase) *ReceiveWindowService {
    return &ReceiveWindowService{
        useCase: useCase,
    }
}

func (s *ReceiveWindowService) GetAllWindowData() ([]entities.WindowSensor, error) {
    data, err := s.useCase.GetAllWindowData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de ventana: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveWindowService) GetWindowDataByID(id int) (*entities.WindowSensor, error) {
    data, err := s.useCase.GetWindowDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de ventana por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveWindowService) GetLatestWindowData() (*entities.WindowSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de ventana")
    }
    return s.latestData, nil
}

func (s *ReceiveWindowService) SerializeWindowData(data []entities.WindowSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de ventana: %v", err)
        return "", err
    }
    return string(jsonData), nil
}