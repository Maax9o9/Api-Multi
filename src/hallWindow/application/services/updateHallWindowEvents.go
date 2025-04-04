package service

import (
    "log"
)

type UpdateHallWindowUseCase interface {
    UpdateStatus(id int, status int) error
}

type UpdateHallWindowService struct {
    useCase UpdateHallWindowUseCase
}

func NewUpdateHallWindowService(useCase UpdateHallWindowUseCase) *UpdateHallWindowService {
    return &UpdateHallWindowService{
        useCase: useCase,
    }
}

func (s *UpdateHallWindowService) UpdateStatus(id int, status int) error {
    log.Printf("Iniciando actualización del estado de HallWindow (ID: %d, Status: %d)", id, status)
    
    err := s.useCase.UpdateStatus(id, status)
    if err != nil {
        log.Printf("Error al actualizar el estado de HallWindow (ID: %d): %v", id, err)
        return err
    }
    
    log.Printf("Estado de HallWindow actualizado correctamente (ID: %d, Status: %d)", id, status)
    return nil
}