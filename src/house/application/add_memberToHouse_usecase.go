package application

import (
	"errors"
	"time"

	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type AddMemberToHouseUseCase struct {
	houseRepo  domain.HouseRepository
	memberRepo domain.HouseMemberRepository
}

func NewAddMemberToHouseUseCase(houseRepo domain.HouseRepository, memberRepo domain.HouseMemberRepository) *AddMemberToHouseUseCase {
	return &AddMemberToHouseUseCase{
		houseRepo:  houseRepo,
		memberRepo: memberRepo,
	}
}

func (uc *AddMemberToHouseUseCase) Execute(houseID, userID int, role string, addedByID int) error {
	// Verificar permisos del que añade
	house, err := uc.houseRepo.GetByID(houseID)
	if err != nil {
		return err
	}

	// El propietario siempre tiene permisos
	if house.OwnerID != addedByID {
		// Verificar si el que añade es administrador
		memberInfo, err := uc.memberRepo.GetByUserAndHouse(addedByID, houseID)
		if err != nil || (memberInfo.Role != "owner" && memberInfo.Role != "admin") {
			return errors.New("no tiene permisos para añadir miembros")
		}
	}

	// Verificar si el usuario ya es miembro
	existingMember, _ := uc.memberRepo.GetByUserAndHouse(userID, houseID)
	if existingMember != nil {
		return errors.New("el usuario ya es miembro de esta casa")
	}

	// Añadir nuevo miembro
	member := &entities.HouseMember{
		HouseID:  houseID,
		UserID:   userID,
		Role:     role,
		JoinedAt: time.Now().Unix(),
	}

	return uc.memberRepo.SaveMember(member)
}
