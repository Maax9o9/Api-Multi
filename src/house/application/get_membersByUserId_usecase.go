package application

import (
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type GetMembersByUserIDUseCase struct {
	memberRepo domain.HouseMemberRepository
}

func NewGetMembersByUserIDUseCase(memberRepo domain.HouseMemberRepository) *GetMembersByUserIDUseCase {
	return &GetMembersByUserIDUseCase{
		memberRepo: memberRepo,
	}
}

func (uc *GetMembersByUserIDUseCase) Execute(userID int) ([]*entities.HouseMember, error) {
	return uc.memberRepo.GetByUserID(userID)
}
