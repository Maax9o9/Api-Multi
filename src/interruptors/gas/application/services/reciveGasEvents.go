package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/gas/domain/entities"
)

type GetAllGasUseCase interface {
    GetAllGasData() ([]entities.GasSensor, error)
}

type GetGasByIDUseCase interface {
    GetGasDataByID(id int) (*entities.GasSensor, error)
}

type ReceiveGasService struct {
    getAllUseCase  GetAllGasUseCase
    getByIDUseCase GetGasByIDUseCase
    latestData     *entities.GasSensor
}

func NewReceiveGasService(getAllUseCase GetAllGasUseCase, getByIDUseCase GetGasByIDUseCase) *ReceiveGasService {
    return &ReceiveGasService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveGasService) GetLatestGasData() (*entities.GasSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de gas")
    }
    return s.latestData, nil
}

func (s *ReceiveGasService) GetAllGasData() ([]entities.GasSensor, error) {
    data, err := s.getAllUseCase.GetAllGasData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de gas: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveGasService) GetGasDataByID(id int) (*entities.GasSensor, error) {
    data, err := s.getByIDUseCase.GetGasDataByID(id)
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