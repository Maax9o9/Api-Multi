package application

import (
    "Multi/src/interruptors/light/domain"
    "Multi/src/interruptors/light/domain/entities"
    "Multi/src/interruptors/light/application/repositorys"
    "encoding/json"
    "log"
)

type AlertLightUseCase struct {
    repo       domain.LightRepository
    rabbitRepo *repositorys.RabbitRepository
}

func NewAlertLightUseCase(repo domain.LightRepository, rabbitRepo *repositorys.RabbitRepository) *AlertLightUseCase {
    return &AlertLightUseCase{
        repo:       repo,
        rabbitRepo: rabbitRepo,
    }
}

func (uc *AlertLightUseCase) CreateLightData(data *entities.LightData) error {
    return uc.repo.Create(data)
}

func (uc *AlertLightUseCase) ProcessLightData(message []byte) error {
    var lightData entities.LightData
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