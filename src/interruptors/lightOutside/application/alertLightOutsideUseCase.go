package application

import (
    "Multi/src/interruptors/lightOutside/domain"
    "Multi/src/interruptors/lightOutside/domain/entities"
    "Multi/src/interruptors/lightOutside/application/repositorys"
    "encoding/json"
    "log"
)

type AlertLightUseCase struct {
    repo       domain.LightOutsideRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertLightUseCase(repo domain.LightOutsideRepository, rabbitRepo *repositorys.RabbitRepository) *AlertLightUseCase {
    return &AlertLightUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertLightUseCase) CreateLightData(data *entities.LightOutsideData) error {
    return uc.repo.Create(data)
}

func (uc *AlertLightUseCase) ProcessLightData(message []byte) error {
    var lightData entities.LightOutsideData
    err := json.Unmarshal(message, &lightData)
    if err != nil {
        log.Printf("Error unmarshalling light data: %v", err)
        return err
    }

    err = uc.CreateLightData(&lightData)
    if err != nil {
        log.Printf("Error saving light data: %v", err)
        return err
    }

    log.Printf("Light data processed and saved: %+v", lightData)
    return nil
}

func (uc *AlertLightUseCase) ProcessLightCommands(processMessage func(body []byte)) error {
    return uc.rabbitRepo.ProcessLightCommands(processMessage)
}