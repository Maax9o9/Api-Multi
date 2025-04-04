package service

import (
    "log"
)

type UpdateHallDoorUseCase interface {
    UpdateStatus(id int, status int) error
}

type UpdateHallDoorService struct {
    useCase UpdateHallDoorUseCase
}

func NewUpdateHallDoorService(useCase UpdateHallDoorUseCase) *UpdateHallDoorService {
    return &UpdateHallDoorService{
        useCase: useCase,
    }
}

func (s *UpdateHallDoorService) UpdateStatus(id int, status int) error {
    err := s.useCase.UpdateStatus(id, status)
    if err != nil {
        log.Printf("Error al actualizar el estado de HallDoor (ID: %d): %v", id, err)
        return err
    }
    log.Printf("Estado de HallDoor actualizado correctamente (ID: %d, Status: %d)", id, status)
    return nil
}