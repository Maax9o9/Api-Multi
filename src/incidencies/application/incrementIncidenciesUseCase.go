package application

import (
    "Multi/src/incidencies/domain"
    "Multi/src/incidencies/domain/entities"
    "Multi/src/incidencies/application/repositorys"
)

type IncrementIncidencyUseCase struct {
    repo       domain.IncidenciesRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewIncrementIncidencyUseCase(repo domain.IncidenciesRepository, rabbitRepo *repositorys.RabbitRepository) *IncrementIncidencyUseCase {
    return &IncrementIncidencyUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *IncrementIncidencyUseCase) IncrementIncidency(typeNotification string) (*entities.Incidency, error) {
    incidency, err := uc.repo.Increment(typeNotification)
    if err != nil {
        return nil, err
    }

    if incidency != nil {
        err = uc.rabbitRepo.PublishIncidency(*incidency)
        if err != nil {
            return nil, err
        }
    }

    return incidency, nil
}

func (uc *IncrementIncidencyUseCase) ProcessIncidenciesData(processMessage func(body []byte)) error {
    return uc.rabbitRepo.ProcessIncidenciesData(processMessage)
}