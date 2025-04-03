package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/light/domain/entities"
)

type GetAllLightUseCase interface {
    GetAllLightData() ([]entities.LightData, error)
}

type GetLightByIDUseCase interface {
    GetLightDataByID(id int) (*entities.LightData, error)
}

type ReceiveLightService struct {
    getAllUseCase  GetAllLightUseCase
    getByIDUseCase GetLightByIDUseCase
    latestData     *entities.LightData
}

func NewReceiveLightService(getAllUseCase GetAllLightUseCase, getByIDUseCase GetLightByIDUseCase) *ReceiveLightService {
    return &ReceiveLightService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveLightService) GetAllLightData() ([]entities.LightData, error) {
    data, err := s.getAllUseCase.GetAllLightData()
    if err != nil {
        log.Printf("Error al obtener todos los datos de la luz: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveLightService) GetLightDataByID(id int) (*entities.LightData, error) {
    data, err := s.getByIDUseCase.GetLightDataByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de la luz por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveLightService) GetLatestLightData() (*entities.LightData, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de la luz")
    }
    return s.latestData, nil
}

func (s *ReceiveLightService) UpdateLatestLightData() error {
    data, err := s.getAllUseCase.GetAllLightData()
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

func (s *ReceiveLightService) SerializeLightData(data []entities.LightData) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de la luz: %v", err)
        return "", err
    }
    return string(jsonData), nil
}