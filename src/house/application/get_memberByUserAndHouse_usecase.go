package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetMemberByUserAndHouseUseCase struct {
	memberRepo domain.HouseMemberRepository
}

func NewGetMemberByUserAndHouseUseCase(memberRepo domain.HouseMemberRepository) *GetMemberByUserAndHouseUseCase {
	return &GetMemberByUserAndHouseUseCase{
		memberRepo: memberRepo,
	}
}

func (uc *GetMemberByUserAndHouseUseCase) Execute(userID, houseID int) (*entities.HouseMember, error) {
	return uc.memberRepo.GetByUserAndHouse(userID, houseID)
}
