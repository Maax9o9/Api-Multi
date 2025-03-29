package infrestructure

import (
    "Multi/src/interruptors/door/application"
    "Multi/src/interruptors/door/infrestructure/adapters"
    "Multi/src/interruptors/door/infrestructure/controllers"
    "log"

    "github.com/streadway/amqp"
)

func InitDoor() *controllers.DoorController {
    conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }

    rabbitAdapter := adapters.NewRabbitAdapter(channel, "DoorQueue")
    doorUseCase := application.NewDoorUseCase(rabbitAdapter)
    doorController := controllers.NewDoorController(doorUseCase)

    return doorController
}