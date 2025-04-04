package application

import (
    "Multi/src/hallWindow/domain"
)

type UpdateHallWindowUseCase struct {
    repo domain.HallWindowRepository
}

func NewUpdateHallWindowUseCase(repo domain.HallWindowRepository) *UpdateHallWindowUseCase {
    return &UpdateHallWindowUseCase{
        repo: repo,
    }
}

func (uc *UpdateHallWindowUseCase) UpdateStatus(id int, status int) error {
    return uc.repo.UpdateStatus(id, status)
}