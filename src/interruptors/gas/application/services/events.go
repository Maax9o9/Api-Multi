package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/gas/domain/entities"
)

type GasUseCase interface {
    GetAllGasData() ([]entities.GasSensor, error)
    GetGasDataByID(id int) (*entities.GasSensor, error)
    CreateGasData(data *entities.GasSensor) error
}

type GasService struct {
    useCase    GasUseCase
    latestData *entities.GasSensor
}

func NewGasService(useCase GasUseCase) *GasService {
    return &GasService{
        useCase: useCase,
    }
}

func (s *GasService) GetLatestGasData() (*entities.GasSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de gas")
    }
    return s.latestData, nil
}

func (s *GasService) GetAllGasData() ([]entities.GasSensor, error) {
    data, err := s.useCase.GetAllGasData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de gas: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *GasService) GetGasDataByID(id int) (*entities.GasSensor, error) {
    data, err := s.useCase.GetGasDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de gas por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *GasService) CreateGasData(data *entities.GasSensor) error {
    err := s.useCase.CreateGasData(data)
    if err != nil {
        log.Printf("Error al crear datos de gas: %v", err)
        return err
    }
    return nil
}

func (s *GasService) SerializeGasData(data []entities.GasSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de gas: %v", err)
        return "", err
    }
    return string(jsonData), nil
}