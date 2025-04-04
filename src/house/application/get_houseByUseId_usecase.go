package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetHousesByUserIDUseCase struct {
	houseRepo domain.HouseRepository
}

func NewGetHousesByUserIDUseCase(houseRepo domain.HouseRepository) *GetHousesByUserIDUseCase {
	return &GetHousesByUserIDUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *GetHousesByUserIDUseCase) Execute(userID int) ([]*entities.HouseProfile, error) {
	return uc.houseRepo.GetHousesByUserID(userID)
}
