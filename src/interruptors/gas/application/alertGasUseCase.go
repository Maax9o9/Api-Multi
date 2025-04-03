package application

import (
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
    "Multi/src/interruptors/gas/application/repositorys"
    "encoding/json"
    "log"
)

type AlertGasUseCase struct {
    repo       domain.GasRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertGasUseCase(repo domain.GasRepository, rabbitRepo *repositorys.RabbitRepository) *AlertGasUseCase {
    return &AlertGasUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertGasUseCase) CreateGasData(data *entities.GasSensor) error {
    return uc.repo.Create(data)
}

func (uc *AlertGasUseCase) ProcessGasData(message []byte) error {
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