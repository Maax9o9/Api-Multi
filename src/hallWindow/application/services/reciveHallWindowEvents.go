package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/hallWindow/domain/entities"
)

type GetAllHallWindowUseCase interface {
    GetAllHallWindows() ([]entities.HallWindow, error)
}

type GetHallWindowByIDUseCase interface {
    GetHallWindowByID(id int) (*entities.HallWindow, error)
}

type ReceiveHallWindowService struct {
    getAllUseCase  GetAllHallWindowUseCase
    getByIDUseCase GetHallWindowByIDUseCase
    latestData     *entities.HallWindow
}

func NewReceiveHallWindowService(getAllUseCase GetAllHallWindowUseCase, getByIDUseCase GetHallWindowByIDUseCase) *ReceiveHallWindowService {
    return &ReceiveHallWindowService{
        getAllUseCase:  getAllUseCase,
        getByIDUseCase: getByIDUseCase,
    }
}

func (s *ReceiveHallWindowService) GetAllHallWindows() ([]entities.HallWindow, error) {
    data, err := s.getAllUseCase.GetAllHallWindows()
    if err != nil {
        log.Printf("Error al obtener todos los datos de HallWindow: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveHallWindowService) GetHallWindowByID(id int) (*entities.HallWindow, error) {
    data, err := s.getByIDUseCase.GetHallWindowByID(id)
    if err != nil {
        log.Printf("Error al obtener datos de HallWindow por ID: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *ReceiveHallWindowService) GetLatestHallWindowData() (*entities.HallWindow, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes")
    }
    return s.latestData, nil
}

func (s *ReceiveHallWindowService) UpdateLatestHallWindowData() error {
    data, err := s.getAllUseCase.GetAllHallWindows()
    if err != nil {
        log.Printf("Error al actualizar los datos de HallWindow: %v", err)
        return err
    }

    if len(data) > 0 {
        s.latestData = &data[len(data)-1]
        log.Printf("Últimos datos de HallWindow actualizados: %+v", s.latestData)
    }
    return nil
}

func (s *ReceiveHallWindowService) SerializeHallWindowData(data []entities.HallWindow) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de HallWindow: %v", err)
        return "", err
    }
    return string(jsonData), nil
}