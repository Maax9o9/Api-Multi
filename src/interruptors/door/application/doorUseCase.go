package application

import (
    "Multi/src/interruptors/door/domain"
    "Multi/src/interruptors/door/domain/entities"
    "Multi/src/interruptors/door/application/repositorys"
    "encoding/json"
    "log"
)

type DoorUseCase struct {
    repo       domain.DoorRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewDoorUseCase(repo domain.DoorRepository, rabbitRepo *repositorys.RabbitRepository) *DoorUseCase {
    return &DoorUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *DoorUseCase) GetAll() ([]entities.DoorData, error) {
    return uc.repo.GetAll()
}

func (uc *DoorUseCase) GetByID(id int) (*entities.DoorData, error) {
    return uc.repo.GetByID(id)
}

func (uc *DoorUseCase) Create(data *entities.DoorData) error {
    return uc.repo.Create(data)
}

func (uc *DoorUseCase) ProcessDoorData(message []byte) error {
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