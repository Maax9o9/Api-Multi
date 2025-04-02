package repositorys

import (
    "Multi/src/interruptors/window/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessWindowCommands(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing window command: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process window commands: %v", err)
        return err
    }
    return nil
}