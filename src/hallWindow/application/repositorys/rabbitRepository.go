package repositorys

import (
    "Multi/src/hallWindow/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessHallWindowData(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing hall window data: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process hall window data: %v", err)
        return err
    }
    return nil
}