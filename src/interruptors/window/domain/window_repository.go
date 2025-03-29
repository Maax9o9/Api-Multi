package domain

import "Multi/src/interruptors/window/domain/entities"

type WindowRepository interface {
    PublishWindowCommand(command entities.WindowCommand) error
}