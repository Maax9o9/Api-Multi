package domain

import "Multi/src/interruptors/window/domain/entities"

type WindowRepository interface {
    Create(data *entities.WindowSensor) error
    GetAll() ([]entities.WindowSensor, error)
    GetByID(id int) (*entities.WindowSensor, error)
}