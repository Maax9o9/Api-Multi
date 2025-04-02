package application

import (
    "Multi/src/interruptors/light/domain"
    "Multi/src/interruptors/light/domain/entities"
    "Multi/src/interruptors/light/application/repositorys"
)

type LightUseCase struct {
    repo       domain.LightRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewLightUseCase(repo domain.LightRepository, rabbitRepo *repositorys.RabbitRepository) *LightUseCase {
    return &LightUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *LightUseCase) GetAllLightData() ([]entities.LightData, error) {
    return uc.repo.GetAll()
}

func (uc *LightUseCase) GetLightDataByID(id int) (*entities.LightData, error) {
    return uc.repo.GetByID(id)
}

func (uc *LightUseCase) CreateLightData(data *entities.LightData) error {
    return uc.repo.Create(data)
}

func (uc *LightUseCase) ProcessLightCommands(processMessage func(body []byte)) error {
    return uc.rabbitRepo.ProcessLightCommands(processMessage)
}