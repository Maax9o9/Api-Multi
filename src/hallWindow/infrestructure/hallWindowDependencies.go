package infrestructure

import (
    "Multi/src/hallWindow/application"
    "Multi/src/hallWindow/application/repositorys"
    "Multi/src/hallWindow/application/services"
    "Multi/src/hallWindow/infrestructure/adapters"
    "log"
)

func InitHallWindow() (*service.AlertHallWindowService, *service.ReceiveHallWindowService, *service.UpdateHallWindowService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "hall.window",  // Queue
        "sensor.hallWindow", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    hallWindowRepo := NewPostgres()

    getAllHallWindowUseCase := application.NewReceiveHallWindowUseCase(hallWindowRepo)
    getHallWindowByIDUseCase := application.NewReceiveHallWindowByIDUseCase(hallWindowRepo)
    alertHallWindowUseCase := application.NewAlertHallWindowUseCase(hallWindowRepo, rabbitRepo)
    updateHallWindowUseCase := application.NewUpdateHallWindowUseCase(hallWindowRepo)

    alertHallWindowService := service.NewAlertHallWindowService(alertHallWindowUseCase)
    receiveHallWindowService := service.NewReceiveHallWindowService(getAllHallWindowUseCase, getHallWindowByIDUseCase)
    updateHallWindowService := service.NewUpdateHallWindowService(updateHallWindowUseCase)

    return alertHallWindowService, receiveHallWindowService, updateHallWindowService, rabbitMQ
}