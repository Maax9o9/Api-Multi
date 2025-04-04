package application

import "Multi/src/house/domain"

type DeleteHouseUseCase struct {
	houseRepo domain.HouseRepository
}

func NewDeleteHouseUseCase(houseRepo domain.HouseRepository) *DeleteHouseUseCase {
	return &DeleteHouseUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *DeleteHouseUseCase) Execute(id int) error {
	return uc.houseRepo.Delete(id)
}
