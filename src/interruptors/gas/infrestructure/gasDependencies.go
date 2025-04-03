package infrestructure

import (
    "Multi/src/interruptors/gas/application"
    "Multi/src/interruptors/gas/application/repositorys"
    "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/infrestructure/adapters"
    "log"
)

func InitGas() (*service.AlertGasService, *service.ReceiveGasService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "gas",          // Queue
        "sensor.mq2",   // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    gasRepo := NewPostgres()

    alertGasUseCase := application.NewAlertGasUseCase(gasRepo, rabbitRepo)
    receiveGasUseCase := application.NewReceiveGasUseCase(gasRepo)

    alertGasService := service.NewAlertGasService(alertGasUseCase)
    receiveGasService := service.NewReceiveGasService(receiveGasUseCase)

    return alertGasService, receiveGasService, rabbitMQ
}