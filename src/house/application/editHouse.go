package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
)

type EditHouseUseCase struct {
    repo domain.HouseRepository
}

func NewEditHouseUseCase(repo domain.HouseRepository) *EditHouseUseCase {
    return &EditHouseUseCase{
        repo: repo,
    }
}

func (uc *EditHouseUseCase) UpdateHouse(house *entities.HouseProfile) error {
    return uc.repo.Update(house)
}