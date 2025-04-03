package application

import (
    "Multi/src/interruptors/window/domain"
    "Multi/src/interruptors/window/domain/entities"
    "Multi/src/interruptors/window/application/repositorys"
    "encoding/json"
    "log"
)

type AlertWindowUseCase struct {
    repo       domain.WindowRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertWindowUseCase(repo domain.WindowRepository, rabbitRepo *repositorys.RabbitRepository) *AlertWindowUseCase {
    return &AlertWindowUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertWindowUseCase) CreateWindowData(data *entities.WindowSensor) error {
    return uc.repo.Create(data)
}

func (uc *AlertWindowUseCase) ProcessWindowData(message []byte) error {
    var windowData entities.WindowSensor
    err := json.Unmarshal(message, &windowData)
    if err != nil {
        log.Printf("Error unmarshalling window data: %v", err)
        return err
    }

    err = uc.CreateWindowData(&windowData)
    if err != nil {
        log.Printf("Error saving window data: %v", err)
        return err
    }

    log.Printf("Window data processed and saved: %+v", windowData)
    return nil
}