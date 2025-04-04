package application

import (
    "Multi/src/hallDoor/domain"
    "Multi/src/hallDoor/domain/entities"
)

type ReceiveHallDoorByIDUseCase struct {
    repo domain.HallDoorRepository
}

func NewReceiveHallDoorByIDUseCase(repo domain.HallDoorRepository) *ReceiveHallDoorByIDUseCase {
    return &ReceiveHallDoorByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveHallDoorByIDUseCase) GetHallDoorByID(id int) (*entities.HallDoor, error) {
    return uc.repo.GetByID(id)
}