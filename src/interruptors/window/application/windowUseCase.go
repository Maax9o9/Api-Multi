package application

import (
    "Multi/src/interruptors/window/domain"
    "Multi/src/interruptors/window/domain/entities"
    "Multi/src/interruptors/window/application/repositorys"
)

type WindowUseCase struct {
    repo       domain.WindowRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewWindowUseCase(repo domain.WindowRepository, rabbitRepo *repositorys.RabbitRepository) *WindowUseCase {
    return &WindowUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *WindowUseCase) GetAllWindowData() ([]entities.WindowSensor, error) {
    return uc.repo.GetAll()
}

func (uc *WindowUseCase) GetWindowDataByID(id int) (*entities.WindowSensor, error) {
    return uc.repo.GetByID(id)
}

func (uc *WindowUseCase) CreateWindowData(data *entities.WindowSensor) error {
    return uc.repo.Create(data)
}