package application

import (
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
    "Multi/src/interruptors/gas/application/repositorys"
)

type GasUseCase struct {
    repo       domain.GasRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewGasUseCase(repo domain.GasRepository, rabbitRepo *repositorys.RabbitRepository) *GasUseCase {
    return &GasUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *GasUseCase) GetAllGasData() ([]entities.GasSensor, error) {
    return uc.repo.GetAll()
}

func (uc *GasUseCase) GetGasDataByID(id int) (*entities.GasSensor, error) {
    return uc.repo.GetByID(id)
}

func (uc *GasUseCase) CreateGasData(data *entities.GasSensor) error {
    return uc.repo.Create(data)
}