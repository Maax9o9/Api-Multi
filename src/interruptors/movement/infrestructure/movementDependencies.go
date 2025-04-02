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
        "amqp://user:password@localhost:5672/",
        "MovementQueue",
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