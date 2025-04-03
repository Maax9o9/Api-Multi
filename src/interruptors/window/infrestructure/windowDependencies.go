package infrestructure

import (
    "Multi/src/interruptors/window/application"
    "Multi/src/interruptors/window/application/repositorys"
    "Multi/src/interruptors/window/application/services"
    "Multi/src/interruptors/window/infrestructure/adapters"
    "log"
)

func InitWindow() (*service.WindowService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",
        "ventana",                          
        "sensor.ventana",                     
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    windowRepo := NewPostgres()

    windowUseCase := application.NewWindowUseCase(windowRepo, rabbitRepo)

    windowService := service.NewWindowService(windowUseCase)

    return windowService, rabbitMQ
}