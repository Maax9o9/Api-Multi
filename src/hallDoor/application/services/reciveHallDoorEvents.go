package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/hallDoor/domain/entities"
)

type GetAllHallDoorUseCase interface {
    GetAllHallDoors() ([]entities.HallDoor, error)
}

type GetHallDoorByIDUseCase interface {
    GetHallDoorByID(id int) (*entities.HallDoor, error)
}

type ReceiveHallDoorService struct {
    getAllUseCase  GetAllHallDoorUseCase
    getByIDUseCase GetHallDoorByIDUseCase
    latestData     *entities.HallDoor
}

func NewReceiveHallDoorService(getAllUseCase GetAllHallDoorUseCase, getByIDUseCase GetHallDoorByIDUseCase) *ReceiveHallDoorService {
    return &ReceiveHallDoorService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveHallDoorService) GetAllHallDoors() ([]entities.HallDoor, error) {
    data, err := s.getAllUseCase.GetAllHallDoors()
    if err != nil {
        log.Printf("Error al obtener todos los datos de HallDoor: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveHallDoorService) GetHallDoorByID(id int) (*entities.HallDoor, error) {
    data, err := s.getByIDUseCase.GetHallDoorByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de HallDoor por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveHallDoorService) GetLatestHallDoorData() (*entities.HallDoor, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes")
    }
    return s.latestData, nil
}

func (s *ReceiveHallDoorService) UpdateLatestHallDoorData() error {
    data, err := s.getAllUseCase.GetAllHallDoors()
    if err != nil {
        log.Printf("Error al actualizar los datos de HallDoor: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1]
        log.Printf("Últimos datos de HallDoor actualizados: %+v", s.latestData)
    }
    return nil
}

func (s *ReceiveHallDoorService) SerializeHallDoorData(data []entities.HallDoor) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de HallDoor: %v", err)
        return "", err
    }
    return string(jsonData), nil
}