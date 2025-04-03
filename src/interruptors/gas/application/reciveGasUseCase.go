package application

import (
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
)

type ReceiveGasUseCase struct {
    repo domain.GasRepository
}

func NewReceiveGasUseCase(repo domain.GasRepository) *ReceiveGasUseCase {
    return &ReceiveGasUseCase{
        repo: repo,
    }
}

func (uc *ReceiveGasUseCase) GetAllGasData() ([]entities.GasSensor, error) {
    return uc.repo.GetAll()
}