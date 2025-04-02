package domain

import "Multi/src/interruptors/light/domain/entities"

type LightRepository interface {
    Create(data *entities.LightData) error
    GetAll() ([]entities.LightData, error)
    GetByID(id int) (*entities.LightData, error)
}