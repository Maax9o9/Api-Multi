package infrestructure

import (
    "Multi/src/interruptors/gas/application"
    "Multi/src/interruptors/gas/application/repositorys"
    "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/infrestructure/adapters"
    "log"
)

func InitGas() (*service.GasService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://user:password@localhost:5672/",
        "GasQueue",
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    gasRepo := NewPostgres()

    gasUseCase := application.NewGasUseCase(gasRepo, rabbitRepo)

    gasService := service.NewGasService(gasUseCase)

    return gasService, rabbitMQ
}