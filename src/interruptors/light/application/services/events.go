package service

import (
    "Multi/src/interruptors/light/domain/entities"
    "fmt"
    "log"
)

type LightUseCase interface {
    GetAllLightData() ([]entities.LightData, error)
    GetLightDataByID(id int) (*entities.LightData, error)
    CreateLightData(data *entities.LightData) error
}

type LightService struct {
    useCase    LightUseCase
    latestData *entities.LightData
}

func NewLightService(useCase LightUseCase) *LightService {
    return &LightService{
        useCase: useCase,
    }
}

func (s *LightService) GetAllLightData() ([]entities.LightData, error) {
    data, err := s.useCase.GetAllLightData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de la luz: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *LightService) GetLightDataByID(id int) (*entities.LightData, error) {
    data, err := s.useCase.GetLightDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de la luz por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *LightService) CreateLightData(data *entities.LightData) error {
    err := s.useCase.CreateLightData(data)
    if err != nil {
        log.Printf("Error al crear datos de la luz: %v", err)
        return err
    }
    return nil
}

func (s *LightService) GetLatestLightData() (*entities.LightData, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de la luz")
    }
    return s.latestData, nil
}

func (s *LightService) UpdateLatestLightData() error {
    data, err := s.useCase.GetAllLightData()
    if err != nil {
        log.Printf("Error al actualizar los datos de la luz: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1]
        log.Printf("Últimos datos de la luz actualizados: %+v", s.latestData)
    }
    return nil
}