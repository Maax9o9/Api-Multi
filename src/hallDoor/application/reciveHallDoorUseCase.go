package application

import (
    "Multi/src/hallDoor/domain"
    "Multi/src/hallDoor/domain/entities"
)

type ReceiveHallDoorUseCase struct {
    repo domain.HallDoorRepository
}

func NewReceiveHallDoorUseCase(repo domain.HallDoorRepository) *ReceiveHallDoorUseCase {
    return &ReceiveHallDoorUseCase{
        repo: repo,
    }
}

func (uc *ReceiveHallDoorUseCase) GetAllHallDoors() ([]entities.HallDoor, error) {
    return uc.repo.GetAll()
}