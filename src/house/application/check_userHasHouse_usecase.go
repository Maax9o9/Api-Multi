package application

import "Multi/src/house/domain"

type CheckUserHasHouseUseCase struct {
	houseRepo domain.HouseRepository
}

func NewCheckUserHasHouseUseCase(houseRepo domain.HouseRepository) *CheckUserHasHouseUseCase {
	return &CheckUserHasHouseUseCase{
		houseRepo: houseRepo,
	}
}

func (uc *CheckUserHasHouseUseCase) Execute(userID int) (bool, error) {
	houses, err := uc.houseRepo.GetHousesByUserID(userID)
	if err != nil {
		return false, err
	}

	return len(houses) > 0, nil
}
