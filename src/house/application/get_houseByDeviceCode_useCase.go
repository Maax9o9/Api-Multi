package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetHouseByDeviceCodeUseCase struct {
	houseRepo domain.HouseRepository
}

func NewGetHouseByDeviceCodeUseCase(houseRepo domain.HouseRepository) *GetHouseByDeviceCodeUseCase {
	return &GetHouseByDeviceCodeUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *GetHouseByDeviceCodeUseCase) Execute(deviceCode string) (*entities.HouseProfile, error) {
	return uc.houseRepo.GetByDeviceCode(deviceCode)
}
