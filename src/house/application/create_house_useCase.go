package application

import (
	"time"

	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type CreateHouseUseCase struct {
	houseRepo  domain.HouseRepository
	memberRepo domain.HouseMemberRepository
}

func NewCreateHouseUseCase(houseRepo domain.HouseRepository, memberRepo domain.HouseMemberRepository) *CreateHouseUseCase {
	return &CreateHouseUseCase{
		houseRepo:  houseRepo,
		memberRepo: memberRepo,
	}
}

func (uc *CreateHouseUseCase) Execute(house *entities.HouseProfile, ownerID int) (int, error) {
	house.OwnerID = ownerID
	house.CreatedAt = time.Now().Unix()

	houseID, err := uc.houseRepo.Save(house)
	if err != nil {
		return 0, err
	}

	// Añadir al creador como miembro propietario
	member := &entities.HouseMember{
		HouseID:  houseID,
		UserID:   ownerID,
		Role:     "owner",
		JoinedAt: time.Now().Unix(),
	}

	err = uc.memberRepo.SaveMember(member)
	if err != nil {
		// En caso de error, intentar revertir la creación de la casa
		_ = uc.houseRepo.Delete(houseID)
		return 0, err
	}

	return houseID, nil
}
