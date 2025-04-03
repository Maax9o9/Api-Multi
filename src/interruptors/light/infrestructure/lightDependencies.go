package infrestructure

import (
    "Multi/src/interruptors/light/application"
    "Multi/src/interruptors/light/application/repositorys"
    "Multi/src/interruptors/light/application/services"
    "Multi/src/interruptors/light/infrestructure/adapters"
    "log"
)

func InitLight() (*service.AlertLightService, *service.ReceiveLightService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "light",        // Queue
        "light.on",     // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    lightRepo := NewPostgres()

    alertLightUseCase := application.NewAlertLightUseCase(lightRepo, rabbitRepo)
    receiveLightUseCase := application.NewReceiveLightUseCase(lightRepo)

    alertLightService := service.NewAlertLightService(alertLightUseCase)
    receiveLightService := service.NewReceiveLightService(receiveLightUseCase)

    return alertLightService, receiveLightService, rabbitMQ
}