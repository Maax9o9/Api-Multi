package application

import (
    "Multi/src/interruptors/door/domain"
    "Multi/src/interruptors/door/domain/entities"
)

type ReceiveDoorUseCase struct {
    repo domain.DoorRepository
}

func NewReceiveDoorUseCase(repo domain.DoorRepository) *ReceiveDoorUseCase {
    return &ReceiveDoorUseCase{
        repo: repo,
    }
}

func (uc *ReceiveDoorUseCase) GetAll() ([]entities.DoorData, error) {
    return uc.repo.GetAll()
}

func (uc *ReceiveDoorUseCase) GetByID(id int) (*entities.DoorData, error) {
    return uc.repo.GetByID(id)
}