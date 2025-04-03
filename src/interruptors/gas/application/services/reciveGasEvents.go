package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/gas/domain/entities"
)

type ReceiveGasUseCase interface {
    GetAllGasData() ([]entities.GasSensor, error)
    GetGasDataByID(id int) (*entities.GasSensor, error)
}

type ReceiveGasService struct {
    useCase    ReceiveGasUseCase
    latestData *entities.GasSensor
}

func NewReceiveGasService(useCase ReceiveGasUseCase) *ReceiveGasService {
    return &ReceiveGasService{
        useCase: useCase,
    }
}

func (s *ReceiveGasService) GetLatestGasData() (*entities.GasSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de gas")
    }
    return s.latestData, nil
}

func (s *ReceiveGasService) GetAllGasData() ([]entities.GasSensor, error) {
    data, err := s.useCase.GetAllGasData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de gas: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveGasService) GetGasDataByID(id int) (*entities.GasSensor, error) {
    data, err := s.useCase.GetGasDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de gas por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveGasService) SerializeGasData(data []entities.GasSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de gas: %v", err)
        return "", err
    }
    return string(jsonData), nil
}