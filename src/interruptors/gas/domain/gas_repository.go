package domain

import "Multi/src/interruptors/gas/domain/entities"

type GasRepository interface {
    Create(data *entities.GasSensor) error
    GetAll() ([]entities.GasSensor, error)
    GetByID(id int) (*entities.GasSensor, error)
}