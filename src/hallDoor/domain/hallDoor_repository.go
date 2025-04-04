package domain

import "Multi/src/hallDoor/domain/entities"

type HallDoorRepository interface {
    GetAll() ([]entities.HallDoor, error)
    GetByID(id int) (*entities.HallDoor, error)
    Create(hallDoor *entities.HallDoor) error
    UpdateStatus(id int, status int) error
}