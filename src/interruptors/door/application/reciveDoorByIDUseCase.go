package application

import (
    "Multi/src/interruptors/door/domain"
    "Multi/src/interruptors/door/domain/entities"
)

type GetDoorByIDUseCase struct {
    repo domain.DoorRepository
}

func NewGetDoorByIDUseCase(repo domain.DoorRepository) *GetDoorByIDUseCase {
    return &GetDoorByIDUseCase{
        repo: repo,
    }
}

func (uc *GetDoorByIDUseCase) GetByID(id int) (*entities.DoorData, error) {
    return uc.repo.GetByID(id)
}