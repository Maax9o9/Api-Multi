package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/movement/domain/entities"
)

type ReceiveMovementUseCase interface {
    GetAll() ([]entities.MotionSensor, error)
    GetByID(id int) (*entities.MotionSensor, error)
}

type ReceiveMovementService struct {
    useCase    ReceiveMovementUseCase
    latestData *entities.MotionSensor
}

func NewReceiveMovementService(useCase ReceiveMovementUseCase) *ReceiveMovementService {
    return &ReceiveMovementService{
        useCase: useCase,
    }
}

func (s *ReceiveMovementService) GetAllMovementData() ([]entities.MotionSensor, error) {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al obtener todos los datos de movimiento: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveMovementService) GetMovementDataByID(id int) (*entities.MotionSensor, error) {
    data, err := s.useCase.GetByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de movimiento por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveMovementService) GetLatestMovementData() (*entities.MotionSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de movimiento")
    }
    return s.latestData, nil
}

func (s *ReceiveMovementService) UpdateLatestMovementData() error {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al actualizar los datos de movimiento: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1]
        log.Printf("Últimos datos de movimiento actualizados: %+v", s.latestData)
    }
    return nil
}

func (s *ReceiveMovementService) SerializeMovementData(data []entities.MotionSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de movimiento: %v", err)
        return "", err
    }
    return string(jsonData), nil
}