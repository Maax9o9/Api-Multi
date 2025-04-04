package application

import "Multi/src/house/domain"

type DeleteHouseMemberUseCase struct {
	memberRepo domain.HouseMemberRepository
}

func NewDeleteHouseMemberUseCase(memberRepo domain.HouseMemberRepository) *DeleteHouseMemberUseCase {
	return &DeleteHouseMemberUseCase{
		memberRepo: memberRepo,
	}
}

func (uc *DeleteHouseMemberUseCase) Execute(userID, houseID int) error {
	return uc.memberRepo.DeleteMember(userID, houseID)
}
