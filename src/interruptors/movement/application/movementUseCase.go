package application

import (
    "Multi/src/interruptors/movement/domain"
    "Multi/src/interruptors/movement/domain/entities"
    "Multi/src/interruptors/movement/application/repositorys"
    "encoding/json"
    "log"
)

type MovementUseCase struct {
    repo       domain.MovementRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewMovementUseCase(repo domain.MovementRepository, rabbitRepo *repositorys.RabbitRepository) *MovementUseCase {
    return &MovementUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *MovementUseCase) GetAll() ([]entities.MotionSensor, error) {
    return uc.repo.GetAll()
}

func (uc *MovementUseCase) GetByID(id int) (*entities.MotionSensor, error) {
    return uc.repo.GetByID(id)
}

func (uc *MovementUseCase) Create(data *entities.MotionSensor) error {
    return uc.repo.Create(data)
}

func (uc *MovementUseCase) ProcessMovementData(message []byte) error {
    var motionData entities.MotionSensor
    err := json.Unmarshal(message, &motionData)
    if err != nil {
        log.Printf("Error unmarshalling movement data: %v", err)
        return err
    }

    err = uc.Create(&motionData)
    if err != nil {
        log.Printf("Error saving movement data: %v", err)
        return err
    }

    log.Printf("Movement data processed and saved: %+v", motionData)
    return nil
}