package application

import (
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
)

type ShowHouseUseCase struct {
    repo domain.HouseRepository
}

func NewShowHouseUseCase(repo domain.HouseRepository) *ShowHouseUseCase {
    return &ShowHouseUseCase{
        repo: repo,
    }
}

func (uc *ShowHouseUseCase) GetAllHouses() ([]entities.HouseProfile, error) {
    return uc.repo.GetAll()
}

func (uc *ShowHouseUseCase) GetHouseByID(id int) (*entities.HouseProfile, error) {
    return uc.repo.GetByID(id)
}