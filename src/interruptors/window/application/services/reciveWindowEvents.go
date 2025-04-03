package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/window/domain/entities"
)

type GetAllWindowUseCase interface {
    GetAllWindowData() ([]entities.WindowSensor, error)
}

type GetWindowByIDUseCase interface {
    GetWindowDataByID(id int) (*entities.WindowSensor, error)
}

type ReceiveWindowService struct {
    getAllUseCase  GetAllWindowUseCase
    getByIDUseCase GetWindowByIDUseCase
    latestData     *entities.WindowSensor
}

func NewReceiveWindowService(getAllUseCase GetAllWindowUseCase, getByIDUseCase GetWindowByIDUseCase) *ReceiveWindowService {
    return &ReceiveWindowService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveWindowService) GetAllWindowData() ([]entities.WindowSensor, error) {
    data, err := s.getAllUseCase.GetAllWindowData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de ventana: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveWindowService) GetWindowDataByID(id int) (*entities.WindowSensor, error) {
    data, err := s.getByIDUseCase.GetWindowDataByID(id)
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