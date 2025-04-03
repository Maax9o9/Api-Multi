package application

import (
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
    "Multi/src/interruptors/gas/application/repositorys"
    "encoding/json"
    "log"
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

func (uc *GasUseCase) ProcessGasData(message []byte) error {
    var gasData entities.GasSensor
    err := json.Unmarshal(message, &gasData)
    if err != nil {
        log.Printf("Error unmarshalling gas data: %v", err)
        return err
    }

    err = uc.CreateGasData(&gasData)
    if err != nil {
        log.Printf("Error saving gas data: %v", err)
        return err
    }

    log.Printf("Gas data processed and saved: %+v", gasData)
    return nil
}