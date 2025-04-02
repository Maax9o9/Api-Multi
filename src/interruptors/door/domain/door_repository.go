package domain

import "Multi/src/interruptors/door/domain/entities"

type DoorRepository interface {
    Create(data *entities.DoorData) error
    GetAll() ([]entities.DoorData, error)
    GetByID(id int) (*entities.DoorData, error)
}