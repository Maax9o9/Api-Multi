package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
)

type AddHouseUseCase struct {
    repo domain.HouseRepository
}

func NewAddHouseUseCase(repo domain.HouseRepository) *AddHouseUseCase {
    return &AddHouseUseCase{
        repo: repo,
    }
}

func (uc *AddHouseUseCase) AddHouse(house *entities.HouseProfile) error {
    return uc.repo.Create(house)
}