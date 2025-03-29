package application

import (
    "Multi/src/interruptors/light/domain/entities"
    "Multi/src/interruptors/light/domain"
)

type LightUseCase struct {
    repo domain.LightRepository
}

func NewLightUseCase(repo domain.LightRepository) *LightUseCase {
    return &LightUseCase{
        repo: repo,
    }
}

func (uc *LightUseCase) SendLightCommand(command entities.LightCommand) error {
    return uc.repo.PublishLightCommand(command)
}