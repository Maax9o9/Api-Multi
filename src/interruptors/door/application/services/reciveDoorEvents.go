package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/door/domain/entities"
)

type ReceiveDoorUseCase interface {
    GetAll() ([]entities.DoorData, error)
    GetByID(id int) (*entities.DoorData, error)
}

type ReceiveDoorService struct {
    useCase    ReceiveDoorUseCase
    latestData *entities.DoorData
}

func NewReceiveDoorService(useCase ReceiveDoorUseCase) *ReceiveDoorService {
    return &ReceiveDoorService{
        useCase: useCase,
    }
}

func (s *ReceiveDoorService) GetAllDoorData() ([]entities.DoorData, error) {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al obtener todos los datos de la puerta: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveDoorService) GetDoorDataByID(id int) (*entities.DoorData, error) {
    data, err := s.useCase.GetByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de la puerta por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveDoorService) GetLatestDoorData() (*entities.DoorData, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de la puerta")
    }
    return s.latestData, nil
}

func (s *ReceiveDoorService) UpdateLatestDoorData() error {
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

func (s *ReceiveDoorService) SerializeDoorData(data []entities.DoorData) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de la puerta: %v", err)
        return "", err
    }
    return string(jsonData), nil
}