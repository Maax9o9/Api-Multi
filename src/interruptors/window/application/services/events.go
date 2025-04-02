package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/window/domain/entities"
)

type WindowUseCase interface {
    GetAllWindowData() ([]entities.WindowSensor, error)
    GetWindowDataByID(id int) (*entities.WindowSensor, error)
    CreateWindowData(data *entities.WindowSensor) error
}

type WindowService struct {
    useCase    WindowUseCase
    latestData *entities.WindowSensor
}

func NewWindowService(useCase WindowUseCase) *WindowService {
    return &WindowService{
        useCase: useCase,
    }
}

func (s *WindowService) GetLatestWindowData() (*entities.WindowSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de ventana")
    }
    return s.latestData, nil
}

func (s *WindowService) GetAllWindowData() ([]entities.WindowSensor, error) {
    data, err := s.useCase.GetAllWindowData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de ventana: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *WindowService) GetWindowDataByID(id int) (*entities.WindowSensor, error) {
    data, err := s.useCase.GetWindowDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de ventana por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *WindowService) CreateWindowData(data *entities.WindowSensor) error {
    err := s.useCase.CreateWindowData(data)
    if err != nil {
        log.Printf("Error al crear datos de ventana: %v", err)
        return err
    }
    return nil
}

func (s *WindowService) SerializeWindowData(data []entities.WindowSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de ventana: %v", err)
        return "", err
    }
    return string(jsonData), nil
}