package application

import (
    "Multi/src/interruptors/window/domain/entities"
    "Multi/src/interruptors/window/domain"
)

type WindowUseCase struct {
    repo domain.WindowRepository
}

func NewWindowUseCase(repo domain.WindowRepository) *WindowUseCase {
    return &WindowUseCase{
        repo: repo,
    }
}

func (uc *WindowUseCase) SendWindowCommand(command entities.WindowCommand) error {
    return uc.repo.PublishWindowCommand(command)
}