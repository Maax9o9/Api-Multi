package application

import (
    "Multi/src/interruptors/movement/domain"
    "Multi/src/interruptors/movement/domain/entities"
)

type ReceiveMovementUseCase struct {
    repo domain.MovementRepository
}

func NewReceiveMovementUseCase(repo domain.MovementRepository) *ReceiveMovementUseCase {
    return &ReceiveMovementUseCase{
        repo: repo,
    }
}

func (uc *ReceiveMovementUseCase) GetAll() ([]entities.MotionSensor, error) {
    return uc.repo.GetAll()
}