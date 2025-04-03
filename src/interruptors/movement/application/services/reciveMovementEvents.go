package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/interruptors/movement/domain/entities"
)

type GetAllMovementUseCase interface {
    GetAll() ([]entities.MotionSensor, error)
}

type GetMovementByIDUseCase interface {
    GetByID(id int) (*entities.MotionSensor, error)
}

type ReceiveMovementService struct {
    getAllUseCase  GetAllMovementUseCase
    getByIDUseCase GetMovementByIDUseCase
    latestData     *entities.MotionSensor
}

func NewReceiveMovementService(getAllUseCase GetAllMovementUseCase, getByIDUseCase GetMovementByIDUseCase) *ReceiveMovementService {
    return &ReceiveMovementService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveMovementService) GetAllMovementData() ([]entities.MotionSensor, error) {
    data, err := s.getAllUseCase.GetAll()
    if err != nil {
        log.Printf("Error al obtener todos los datos de movimiento: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveMovementService) GetMovementDataByID(id int) (*entities.MotionSensor, error) {
    data, err := s.getByIDUseCase.GetByID(id)
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
    data, err := s.getAllUseCase.GetAll()
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