package infrestructure

import (
    "Multi/src/interruptors/light/application"
    "Multi/src/interruptors/light/infrestructure/adapters"
    "Multi/src/interruptors/light/infrestructure/controllers"
    "log"

    "github.com/streadway/amqp"
)

func InitLight() *controllers.LightController {
    conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }

    rabbitAdapter := adapters.NewRabbitAdapter(channel, "LightQueue")
    lightUseCase := application.NewLightUseCase(rabbitAdapter)
    lightController := controllers.NewLightController(lightUseCase)

    return lightController
}