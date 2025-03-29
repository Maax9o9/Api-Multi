package domain

import "Multi/src/interruptors/movement/domain/entities"

type MovementRepository interface {
    PublishMovementCommand(command entities.MovementCommand) error
}