package domain

import "Multi/src/interruptors/door/domain/entities"

type DoorRepository interface {
    PublishDoorCommand(command entities.DoorCommand) error
}