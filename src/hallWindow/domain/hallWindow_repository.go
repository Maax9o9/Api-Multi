package domain

import "Multi/src/hallWindow/domain/entities"

type HallWindowRepository interface {
    GetAll() ([]entities.HallWindow, error)
    GetByID(id int) (*entities.HallWindow, error)
    Create(hallDoor *entities.HallWindow) error
    UpdateStatus(id int, status int) error
}