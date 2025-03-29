package infrestructure

import (
    "Multi/src/interruptors/movement/application"
    "Multi/src/interruptors/movement/infrestructure/adapters"
    "Multi/src/interruptors/movement/infrestructure/controllers"
    "log"

    "github.com/streadway/amqp"
)

func InitMovement() *controllers.MovementController {
    conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }

    rabbitAdapter := adapters.NewRabbitAdapter(channel, "MovementQueue")
    movementUseCase := application.NewMovementUseCase(rabbitAdapter)
    movementController := controllers.NewMovementController(movementUseCase)

    return movementController
}