package repositorys

import (
    "Multi/src/interruptors/door/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessDoorCommands(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing door command: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process door commands: %v", err)
        return err
    }
    return nil
}