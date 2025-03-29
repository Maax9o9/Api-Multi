package domain

import "Multi/src/house/domain/entities"

type HouseRepository interface {
    Create(house *entities.HouseProfile) error
    GetAll() ([]entities.HouseProfile, error)
    GetByID(id int) (*entities.HouseProfile, error)
    Update(house *entities.HouseProfile) error
}