package repositorys

import (
    "Multi/src/interruptors/gas/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessGasCommands(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing gas command: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process gas commands: %v", err)
        return err
    }
    return nil
}