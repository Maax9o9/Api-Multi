package application

import (
    "Multi/src/interruptors/door/domain"
    "Multi/src/interruptors/door/domain/entities"
    "Multi/src/interruptors/door/application/repositorys"
    "encoding/json"
    "log"
)

type AlertDoorUseCase struct {
    repo       domain.DoorRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertDoorUseCase(repo domain.DoorRepository, rabbitRepo *repositorys.RabbitRepository) *AlertDoorUseCase {
    return &AlertDoorUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertDoorUseCase) Create(data *entities.DoorData) error {
    return uc.repo.Create(data)
}

func (uc *AlertDoorUseCase) ProcessDoorData(message []byte) error {
    var doorData entities.DoorData
    err := json.Unmarshal(message, &doorData)
    if err != nil {
        log.Printf("Error unmarshalling door data: %v", err)
        return err
    }

    err = uc.repo.Create(&doorData)
    if err != nil {
        log.Printf("Error saving door data: %v", err)
        return err
    }

    log.Printf("Door data processed and saved: %+v", doorData)
    return nil
}