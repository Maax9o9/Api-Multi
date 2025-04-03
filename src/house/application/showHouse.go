package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
)

type ShowAllHousesUseCase struct {
    repo domain.HouseRepository
}

func NewShowAllHousesUseCase(repo domain.HouseRepository) *ShowAllHousesUseCase {
    return &ShowAllHousesUseCase{
        repo: repo,
    }
}

func (uc *ShowAllHousesUseCase) GetAllHouses() ([]entities.HouseProfile, error) {
    return uc.repo.GetAll()
}