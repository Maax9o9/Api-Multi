package application

import (
    "Multi/src/incidencies/domain"
    "Multi/src/incidencies/domain/entities"
    "Multi/src/incidencies/application/repositorys"
)

type GetIncidenciesUseCase struct {
    repo       domain.IncidenciesRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewGetIncidenciesUseCase(repo domain.IncidenciesRepository, rabbitRepo *repositorys.RabbitRepository) *GetIncidenciesUseCase {
    return &GetIncidenciesUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *GetIncidenciesUseCase) GetAllIncidencies() ([]entities.Incidency, error) {
    return uc.repo.GetAll()
}

func (uc *GetIncidenciesUseCase) GetIncidencyByType(typeNotification string) (*entities.Incidency, error) {
    return uc.repo.GetByType(typeNotification)
}

func (uc *GetIncidenciesUseCase) ProcessIncidenciesData(processMessage func(body []byte)) error {
    return uc.rabbitRepo.ProcessIncidenciesData(processMessage)
}