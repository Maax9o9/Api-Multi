package infrestructure

import (
    "Multi/src/interruptors/window/application"
    "Multi/src/interruptors/window/application/repositorys"
    "Multi/src/interruptors/window/application/services"
    "Multi/src/interruptors/window/infrestructure/adapters"
    "log"
)

func InitWindow() (*service.AlertWindowService, *service.ReceiveWindowService, *adapters.RabbitConsumer) {
    // Inicializar RabbitMQ
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "ventana",      // Queue
        "sensor.ventana", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    windowRepo := NewPostgres()

    alertWindowUseCase := application.NewAlertWindowUseCase(windowRepo, rabbitRepo)
    receiveWindowUseCase := application.NewReceiveWindowUseCase(windowRepo)

    alertWindowService := service.NewAlertWindowService(alertWindowUseCase)
    receiveWindowService := service.NewReceiveWindowService(receiveWindowUseCase)

    return alertWindowService, receiveWindowService, rabbitMQ
}