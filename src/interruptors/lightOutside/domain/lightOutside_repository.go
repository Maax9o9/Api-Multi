package domain

import "Multi/src/interruptors/lightOutside/domain/entities"

type LightOutsideRepository interface {
    Create(data *entities.LightOutsideData) error
    GetAll() ([]entities.LightOutsideData, error)
    GetByID(id int) (*entities.LightOutsideData, error)
}