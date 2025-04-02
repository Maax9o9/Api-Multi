package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/door/domain/entities"
)

type DoorUseCase interface {
    GetAll() ([]entities.DoorData, error)
    GetByID(id int) (*entities.DoorData, error)
    Create(data *entities.DoorData) error
}

type DoorService struct {
    useCase    DoorUseCase
    latestData *entities.DoorData
}

func NewDoorService(useCase DoorUseCase) *DoorService {
    return &DoorService{
        useCase: useCase,
    }
}

func (s *DoorService) GetLatestDoorData() (*entities.DoorData, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de la puerta")
    }
    return s.latestData, nil
}

func (s *DoorService) GetAllDoorData() ([]entities.DoorData, error) {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al obtener todos los datos de la puerta: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *DoorService) GetDoorDataByID(id int) (*entities.DoorData, error) {
    data, err := s.useCase.GetByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de la puerta por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *DoorService) CreateDoorData(data *entities.DoorData) error {
    err := s.useCase.Create(data)
    if err != nil {
        log.Printf("Error al crear datos de la puerta: %v", err)
        return err
    }
    return nil
}

func (s *DoorService) UpdateLatestDoorData() error {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al actualizar los datos de la puerta: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1]
        log.Printf("Últimos datos de la puerta actualizados: %+v", s.latestData)
    }
    return nil
}

func (s *DoorService) SerializeDoorData(data []entities.DoorData) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de la puerta: %v", err)
        return "", err
    }
    return string(jsonData), nil
}