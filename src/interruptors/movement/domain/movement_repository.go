package domain

import "Multi/src/interruptors/movement/domain/entities"

type MovementRepository interface {
    Create(data *entities.MotionSensor) error
    GetAll() ([]entities.MotionSensor, error)
    GetByID(id int) (*entities.MotionSensor, error)
}