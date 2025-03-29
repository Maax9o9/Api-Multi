package infrestructure

import (
    "Multi/src/interruptors/window/application"
    "Multi/src/interruptors/window/infrestructure/adapters"
    "Multi/src/interruptors/window/infrestructure/controllers"
    "log"

    "github.com/streadway/amqp"
)

func InitWindow() *controllers.WindowController {
    conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }

    rabbitAdapter := adapters.NewRabbitAdapter(channel, "WindowQueue")
    windowUseCase := application.NewWindowUseCase(rabbitAdapter)
    windowController := controllers.NewWindowController(windowUseCase)

    return windowController
}