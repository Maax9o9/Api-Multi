package application

import (
    "Multi/src/hallDoor/domain"
)

type UpdateHallDoorUseCase struct {
    repo domain.HallDoorRepository
}

func NewUpdateHallDoorUseCase(repo domain.HallDoorRepository) *UpdateHallDoorUseCase {
    return &UpdateHallDoorUseCase{
        repo: repo,
    }
}

func (uc *UpdateHallDoorUseCase) UpdateStatus(id int, status int) error {
    return uc.repo.UpdateStatus(id, status)
}