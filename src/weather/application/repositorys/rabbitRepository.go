package repositorys

import (
    "Multi/src/weather/infrestructure/adapters"
    "log"
)

type RabbitRepository struct {
    rabbitMQ *adapters.RabbitConsumer
}

func NewRabbitRepository(rabbitMQ *adapters.RabbitConsumer) *RabbitRepository {
    return &RabbitRepository{rabbitMQ: rabbitMQ}
}

func (repo *RabbitRepository) ProcessWeatherData(processMessage func(body []byte)) error {
    err := repo.rabbitMQ.ConsumeMessages(func(body []byte) {
        log.Printf("Processing weather data: %s", body)
        processMessage(body)
    })
    if err != nil {
        log.Printf("Failed to process weather data: %v", err)
        return err
    }
    return nil
}