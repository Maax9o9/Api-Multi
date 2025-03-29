package domain

import "Multi/src/interruptors/light/domain/entities"

type LightRepository interface {
    PublishLightCommand(command entities.LightCommand) error
}