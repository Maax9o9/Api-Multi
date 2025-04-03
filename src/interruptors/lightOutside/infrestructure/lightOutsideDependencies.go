package infrestructure

import (
    "Multi/src/interruptors/lightOutside/application"
    "Multi/src/interruptors/lightOutside/application/repositorys"
    "Multi/src/interruptors/lightOutside/application/services"
    "Multi/src/interruptors/lightOutside/infrestructure/adapters"
    "log"
)

func InitLightOutside() (*service.AlertLightService, *service.ReceiveLightService, *adapters.RabbitConsumer) {
    rabbitMQOutside, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "lightOut", // Queue
        "light.out", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ for lightOutside: %v", err)
    }

    rabbitRepoOutside := repositorys.NewRabbitRepository(rabbitMQOutside)
    lightOutsideRepo := NewPostgres()

    getAllLightOutsideUseCase := application.NewReceiveLightUseCase(lightOutsideRepo)
    getLightOutsideByIDUseCase := application.NewReceiveLightByIDUseCase(lightOutsideRepo)

    alertLightOutsideUseCase := application.NewAlertLightUseCase(lightOutsideRepo, rabbitRepoOutside)
    alertLightOutsideService := service.NewAlertLightService(alertLightOutsideUseCase)
    receiveLightOutsideService := service.NewReceiveLightService(getAllLightOutsideUseCase, getLightOutsideByIDUseCase)

    return alertLightOutsideService, receiveLightOutsideService, rabbitMQOutside
}