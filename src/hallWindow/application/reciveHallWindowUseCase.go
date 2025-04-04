package application

import (
    "Multi/src/hallWindow/domain"
    "Multi/src/hallWindow/domain/entities"
)

type ReceiveHallWindowUseCase struct {
    repo domain.HallWindowRepository
}

func NewReceiveHallWindowUseCase(repo domain.HallWindowRepository) *ReceiveHallWindowUseCase {
    return &ReceiveHallWindowUseCase{
        repo: repo,
    }
}

func (uc *ReceiveHallWindowUseCase) GetAllHallWindows() ([]entities.HallWindow, error) {
    return uc.repo.GetAll()
}