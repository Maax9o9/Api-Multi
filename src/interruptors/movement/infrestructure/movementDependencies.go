package infrestructure

import (
    "Multi/src/interruptors/movement/application"
    "Multi/src/interruptors/movement/application/repositorys"
    "Multi/src/interruptors/movement/application/services"
    "Multi/src/interruptors/movement/infrestructure/adapters"
    "log"
)

func InitMovement() (*service.MovementService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",                     
        "movimiento",                       
        "sensor.pir",                   
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    movementRepo := NewPostgres()

    movementUseCase := application.NewMovementUseCase(movementRepo, rabbitRepo)

    movementService := service.NewMovementService(movementUseCase)

    return movementService, rabbitMQ
}