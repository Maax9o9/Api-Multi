package application

import (
    "Multi/src/interruptors/door/domain/entities"
    "Multi/src/interruptors/door/domain"
)

type DoorUseCase struct {
    repo domain.DoorRepository
}

func NewDoorUseCase(repo domain.DoorRepository) *DoorUseCase {
    return &DoorUseCase{
        repo: repo,
    }
}

func (uc *DoorUseCase) SendDoorCommand(command entities.DoorCommand) error {
    return uc.repo.PublishDoorCommand(command)
}