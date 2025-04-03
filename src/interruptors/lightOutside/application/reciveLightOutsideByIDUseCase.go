package application

import (
    "Multi/src/interruptors/lightOutside/domain"
    "Multi/src/interruptors/lightOutside/domain/entities"
)

type ReceiveLightByIDUseCase struct {
    repo domain.LightOutsideRepository
}

func NewReceiveLightByIDUseCase(repo domain.LightOutsideRepository) *ReceiveLightByIDUseCase {
    return &ReceiveLightByIDUseCase{
        repo: repo,
    }
}

func (uc *ReceiveLightByIDUseCase) GetLightDataByID(id int) (*entities.LightOutsideData, error) {
    return uc.repo.GetByID(id)
}