package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/movement/domain/entities"
)

type MovementUseCase interface {
    GetAll() ([]entities.MotionSensor, error)
    GetByID(id int) (*entities.MotionSensor, error)
    Create(data *entities.MotionSensor) error
}

type MovementService struct {
    useCase    MovementUseCase
    latestData *entities.MotionSensor
}

func NewMovementService(useCase MovementUseCase) *MovementService {
    return &MovementService{
        useCase: useCase,
    }
}

func (s *MovementService) GetLatestMovementData() (*entities.MotionSensor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de movimiento")
    }
    return s.latestData, nil
}

func (s *MovementService) GetAllMovementData() ([]entities.MotionSensor, error) {
    data, err := s.useCase.GetAll()
    if err != nil {
        log.Printf("Error al obtener todos los datos de movimiento: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *MovementService) GetMovementDataByID(id int) (*entities.MotionSensor, error) {
    data, err := s.useCase.GetByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de movimiento por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *MovementService) CreateMovementData(data *entities.MotionSensor) error {
    err := s.useCase.Create(data)
    if err != nil {
        log.Printf("Error al crear datos de movimiento: %v", err)
        return err
    }
    return nil
}

func (s *MovementService) UpdateLatestMovementData() error {
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

func (s *MovementService) SerializeMovementData(data []entities.MotionSensor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de movimiento: %v", err)
        return "", err
    }
    return string(jsonData), nil
}