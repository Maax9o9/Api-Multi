package application

import (
    "Multi/src/hallWindow/domain"
    "Multi/src/hallWindow/domain/entities"
)

type ReceiveHallWindowByIDUseCase struct {
    repo domain.HallWindowRepository
}

func NewReceiveHallWindowByIDUseCase(repo domain.HallWindowRepository) *ReceiveHallWindowByIDUseCase {
    return &ReceiveHallWindowByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveHallWindowByIDUseCase) GetHallWindowByID(id int) (*entities.HallWindow, error) {
    return uc.repo.GetByID(id)
}