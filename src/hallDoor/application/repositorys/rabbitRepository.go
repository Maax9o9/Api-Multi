package repositorys

import (
    "Multi/src/hallDoor/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessHallDoorData(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing hall door data: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process hall door data: %v", err)
        return err
    }
    return nil
}