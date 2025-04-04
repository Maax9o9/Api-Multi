package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetHouseByIDUseCase struct {
	houseRepo domain.HouseRepository
}

func NewGetHouseByIDUseCase(houseRepo domain.HouseRepository) *GetHouseByIDUseCase {
	return &GetHouseByIDUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *GetHouseByIDUseCase) Execute(id int) (*entities.HouseProfile, error) {
	return uc.houseRepo.GetByID(id)
}
