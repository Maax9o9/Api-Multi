package application

import (
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
)

type ReceiveGasByIDUseCase struct {
    repo domain.GasRepository
}

func NewReceiveGasByIDUseCase(repo domain.GasRepository) *ReceiveGasByIDUseCase {
    return &ReceiveGasByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveGasByIDUseCase) GetGasDataByID(id int) (*entities.GasSensor, error) {
    return uc.repo.GetByID(id)
}