package infrestructure

import (
    "Multi/src/interruptors/movement/application"
    "Multi/src/interruptors/movement/application/repositorys"
    "Multi/src/interruptors/movement/application/services"
    "Multi/src/interruptors/movement/infrestructure/adapters"
    "log"
)

func InitMovement() (*service.AlertMovementService, *service.ReceiveMovementService, *adapters.RabbitConsumer) {
    // Inicializar RabbitMQ
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "movimiento",   // Queue
        "sensor.pir",   // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    movementRepo := NewPostgres()

    alertMovementUseCase := application.NewAlertMovementUseCase(movementRepo, rabbitRepo)
    receiveMovementUseCase := application.NewReceiveMovementUseCase(movementRepo)

    alertMovementService := service.NewAlertMovementService(alertMovementUseCase)
    receiveMovementService := service.NewReceiveMovementService(receiveMovementUseCase)

    return alertMovementService, receiveMovementService, rabbitMQ
}