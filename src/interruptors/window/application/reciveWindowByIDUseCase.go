package application

import (
    "Multi/src/interruptors/window/domain"
    "Multi/src/interruptors/window/domain/entities"
)

type ReceiveWindowByIDUseCase struct {
    repo domain.WindowRepository
}

func NewReceiveWindowByIDUseCase(repo domain.WindowRepository) *ReceiveWindowByIDUseCase {
    return &ReceiveWindowByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveWindowByIDUseCase) GetWindowDataByID(id int) (*entities.WindowSensor, error) {
    return uc.repo.GetByID(id)
}