package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
)

type ShowHouseByIDUseCase struct {
    repo domain.HouseRepository
}

func NewShowHouseByIDUseCase(repo domain.HouseRepository) *ShowHouseByIDUseCase {
    return &ShowHouseByIDUseCase{
        repo: repo,
    }
}

func (uc *ShowHouseByIDUseCase) GetHouseByID(id int) (*entities.HouseProfile, error) {
    return uc.repo.GetByID(id)
}