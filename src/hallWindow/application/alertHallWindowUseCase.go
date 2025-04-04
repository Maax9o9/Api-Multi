package application

import (
    "Multi/src/hallWindow/domain"
    "Multi/src/hallWindow/domain/entities"
    "Multi/src/hallWindow/application/repositorys"
    "encoding/json"
    "log"
)

type AlertHallWindowUseCase struct {
    repo       domain.HallWindowRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertHallWindowUseCase(repo domain.HallWindowRepository, rabbitRepo *repositorys.RabbitRepository) *AlertHallWindowUseCase {
    return &AlertHallWindowUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertHallWindowUseCase) CreateHallWindow(data *entities.HallWindow) error {
    return uc.repo.Create(data)
}

func (uc *AlertHallWindowUseCase) ProcessHallWindowData(message []byte) error {
    var hallWindowData entities.HallWindow
    err := json.Unmarshal(message, &hallWindowData)
    if err != nil {
        log.Printf("Error unmarshalling hall window data: %v", err)
        return err
    }

    err = uc.CreateHallWindow(&hallWindowData)
    if err != nil {
        log.Printf("Error saving hall window data: %v", err)
        return err
    }

    log.Printf("Hall window data processed and saved: %+v", hallWindowData)
    return nil
}