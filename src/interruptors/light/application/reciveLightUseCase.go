package application

import (
    "Multi/src/interruptors/light/domain"
    "Multi/src/interruptors/light/domain/entities"
)

type ReceiveLightUseCase struct {
    repo domain.LightRepository
}

func NewReceiveLightUseCase(repo domain.LightRepository) *ReceiveLightUseCase {
    return &ReceiveLightUseCase{
        repo: repo,
    }
}

func (uc *ReceiveLightUseCase) GetAllLightData() ([]entities.LightData, error) {
    return uc.repo.GetAll()
}

func (uc *ReceiveLightUseCase) GetLightDataByID(id int) (*entities.LightData, error) {
    return uc.repo.GetByID(id)
}