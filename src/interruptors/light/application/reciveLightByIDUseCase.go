package application

import (
    "Multi/src/interruptors/light/domain"
    "Multi/src/interruptors/light/domain/entities"
)

type ReceiveLightByIDUseCase struct {
    repo domain.LightRepository
}

func NewReceiveLightByIDUseCase(repo domain.LightRepository) *ReceiveLightByIDUseCase {
    return &ReceiveLightByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveLightByIDUseCase) GetLightDataByID(id int) (*entities.LightData, error) {
    return uc.repo.GetByID(id)
}