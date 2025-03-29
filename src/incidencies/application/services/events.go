package service

import (
    "encoding/json"
    "fmt"
    "log"
    "Multi/src/incidencies/domain/entities"
)

type GetIncidenciesUseCase interface {
    GetAllIncidencies() ([]entities.Incidency, error)
    GetIncidencyByType(typeNotification string) (*entities.Incidency, error)
}

type IncrementIncidencyUseCase interface {
    IncrementIncidency(typeNotification string) (*entities.Incidency, error)
}

type IncidenciesService struct {
    getUseCase       GetIncidenciesUseCase
    incrementUseCase IncrementIncidencyUseCase
    latestData       *entities.Incidency
}

func NewIncidenciesService(
    getUseCase GetIncidenciesUseCase,
    incrementUseCase IncrementIncidencyUseCase,
) *IncidenciesService {
    return &IncidenciesService{
        getUseCase:       getUseCase,
        incrementUseCase: incrementUseCase,
    }
}

func (s *IncidenciesService) GetLatestIncidency() (*entities.Incidency, error) {
    if s.latestData == nil {
        return nil, fmt.Errorf("no hay datos recientes de incidencias")
    }
    return s.latestData, nil
}

func (s *IncidenciesService) GetAllIncidencies() ([]entities.Incidency, error) {
    data, err := s.getUseCase.GetAllIncidencies()
    if err != nil {
        log.Printf("Error al obtener todas las incidencias: %v", err)
        return nil, err
    }
    return data, nil
}

func (s *IncidenciesService) UpdateLatestIncidency(typeNotification string) error {
    incidency, err := s.getUseCase.GetIncidencyByType(typeNotification)
    if err != nil {
        log.Printf("Error al actualizar la última incidencia: %v", err)
        return err
    }

    s.latestData = incidency
    log.Printf("Última incidencia actualizada: %+v", s.latestData)
    return nil
}

func (s *IncidenciesService) IncrementIncidency(typeNotification string) error {
    _, err := s.incrementUseCase.IncrementIncidency(typeNotification)
    if err != nil {
        log.Printf("Error al incrementar la incidencia: %v", err)
        return err
    }
    return nil
}

func (s *IncidenciesService) SerializeIncidenciesData(data []entities.Incidency) (string, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error al serializar los datos de incidencias: %v", err)
        return "", err
    }
    return string(jsonData), nil
}