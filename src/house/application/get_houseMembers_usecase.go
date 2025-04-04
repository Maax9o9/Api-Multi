package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetHouseMembersUseCase struct {
	memberRepo domain.HouseMemberRepository
}

func NewGetHouseMembersUseCase(memberRepo domain.HouseMemberRepository) *GetHouseMembersUseCase {
	return &GetHouseMembersUseCase{
		memberRepo: memberRepo,
	}
}

func (uc *GetHouseMembersUseCase) Execute(houseID int) ([]*entities.HouseMember, error) {
	return uc.memberRepo.GetHouseMembers(houseID)
}
