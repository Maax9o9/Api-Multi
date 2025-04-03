package application

import (
    "Multi/src/interruptors/window/domain"
    "Multi/src/interruptors/window/domain/entities"
)

type ReceiveWindowUseCase struct {
    repo domain.WindowRepository
}

func NewReceiveWindowUseCase(repo domain.WindowRepository) *ReceiveWindowUseCase {
    return &ReceiveWindowUseCase{
        repo: repo,
    }
}

func (uc *ReceiveWindowUseCase) GetAllWindowData() ([]entities.WindowSensor, error) {
    return uc.repo.GetAll()
}