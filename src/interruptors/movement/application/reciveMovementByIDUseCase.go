package application

import (
    "Multi/src/interruptors/movement/domain"
    "Multi/src/interruptors/movement/domain/entities"
)

type ReceiveMovementByIDUseCase struct {
    repo domain.MovementRepository
}

func NewReceiveMovementByIDUseCase(repo domain.MovementRepository) *ReceiveMovementByIDUseCase {
    return &ReceiveMovementByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveMovementByIDUseCase) GetByID(id int) (*entities.MotionSensor, error) {
    return uc.repo.GetByID(id)
}