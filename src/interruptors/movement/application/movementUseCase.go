package application

import (
    "Multi/src/interruptors/movement/domain/entities"
    "Multi/src/interruptors/movement/domain"
)

type MovementUseCase struct {
    repo domain.MovementRepository
}

func NewMovementUseCase(repo domain.MovementRepository) *MovementUseCase {
    return &MovementUseCase{
        repo: repo,
    }
}

func (uc *MovementUseCase) SendMovementCommand(command entities.MovementCommand) error {
    return uc.repo.PublishMovementCommand(command)
}