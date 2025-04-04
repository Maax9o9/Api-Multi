package application

import (
    "Multi/src/hallDoor/domain"
    "Multi/src/hallDoor/domain/entities"
    "Multi/src/hallDoor/application/repositorys"
    "encoding/json"
    "log"
)

type AlertHallDoorUseCase struct {
    repo       domain.HallDoorRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertHallDoorUseCase(repo domain.HallDoorRepository, rabbitRepo *repositorys.RabbitRepository) *AlertHallDoorUseCase {
    return &AlertHallDoorUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertHallDoorUseCase) CreateHallDoor(data *entities.HallDoor) error {
    return uc.repo.Create(data)
}

func (uc *AlertHallDoorUseCase) ProcessHallDoorData(message []byte) error {
    var hallDoorData entities.HallDoor
    err := json.Unmarshal(message, &hallDoorData)
    if err != nil {
        log.Printf("Error unmarshalling hall door data: %v", err)
        return err
    }

    err = uc.CreateHallDoor(&hallDoorData)
    if err != nil {
        log.Printf("Error saving hall door data: %v", err)
        return err
    }

    log.Printf("Hall door data processed and saved: %+v", hallDoorData)
    return nil
}