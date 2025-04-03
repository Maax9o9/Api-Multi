package infrestructure

import (
    "Multi/src/interruptors/light/application"
    "Multi/src/interruptors/light/application/repositorys"
    "Multi/src/interruptors/light/application/services"
    "Multi/src/interruptors/light/infrestructure/adapters"
    "log"
)

func InitLight() (*service.LightService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",                       
        "light",                           
        "light.on",                      
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    lightRepo := NewPostgres()

    lightUseCase := application.NewLightUseCase(lightRepo, rabbitRepo)

    lightService := service.NewLightService(lightUseCase)

    return lightService, rabbitMQ
}