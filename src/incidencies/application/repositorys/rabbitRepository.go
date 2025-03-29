package repositorys

import (
    "Multi/src/incidencies/domain/entities"
    "Multi/src/incidencies/infrestructure/adapters"
    "encoding/json"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitAdapter
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitAdapter) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) PublishIncidency(incidency entities.Incidency) error {
    message, err := json.Marshal(incidency)
    if err != nil {
        log.Printf("Error al serializar la incidencia: %v", err)
        return err
    }

    err = repo.rabbitMQ.PublishMessage(message)
    if err != nil {
        log.Printf("Error al publicar la incidencia en RabbitMQ: %v", err)
        return err
    }

    log.Printf("Incidencia publicada en RabbitMQ: %+v", incidency)
    return nil
}

func (repo *RabbitRepository) ProcessIncidenciesData(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing incidency data: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process incidency data: %v", err)
        return err
    }
    return nil
}