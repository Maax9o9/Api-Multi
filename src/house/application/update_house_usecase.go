package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type UpdateHouseUseCase struct {
	houseRepo domain.HouseRepository
}

func NewUpdateHouseUseCase(houseRepo domain.HouseRepository) *UpdateHouseUseCase {
	return &UpdateHouseUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *UpdateHouseUseCase) Execute(house *entities.HouseProfile) error {
	return uc.houseRepo.Update(house)
}
