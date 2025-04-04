package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type SaveHouseMemberUseCase struct {
	memberRepo domain.HouseMemberRepository
}

func NewSaveHouseMemberUseCase(memberRepo domain.HouseMemberRepository) *SaveHouseMemberUseCase {
	return &SaveHouseMemberUseCase{
		memberRepo: memberRepo,
	}
}

func (uc *SaveHouseMemberUseCase) Execute(member *entities.HouseMember) error {
	return uc.memberRepo.SaveMember(member)
}
