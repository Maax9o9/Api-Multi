package application

import (
    "Multi/src/interruptors/lightOutside/domain"
    "Multi/src/interruptors/lightOutside/domain/entities"
)

type ReceiveLightUseCase struct {
    repo domain.LightOutsideRepository
}

func NewReceiveLightUseCase(repo domain.LightOutsideRepository) *ReceiveLightUseCase {
    return &ReceiveLightUseCase{
        repo: repo,
    }
}

func (uc *ReceiveLightUseCase) GetAllLightData() ([]entities.LightOutsideData, error) {
    return uc.repo.GetAll()
}