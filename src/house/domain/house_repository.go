package domain

import (
	"Multi/src/house/domain/entities"
	"io"
)

type HouseRepository interface {
	Save(house *entities.HouseProfile) (int, error)
	GetByID(id int) (*entities.HouseProfile, error)
	GetByDeviceCode(deviceCode string) (*entities.HouseProfile, error)
	GetHousesByUserID(userID int) ([]*entities.HouseProfile, error)
	Update(house *entities.HouseProfile) error
	Delete(id int) error
}

type HouseMemberRepository interface {
	SaveMember(member *entities.HouseMember) error
	GetByUserID(userID int) ([]*entities.HouseMember, error)
	GetByUserAndHouse(userID, houseID int) (*entities.HouseMember, error)
	GetHouseMembers(houseID int) ([]*entities.HouseMember, error)
	DeleteMember(userID, houseID int) error
}

type FileStorage interface {
	SaveFile(filename string, file io.Reader) (string, error)
	DeleteFile(filepath string) error
}
