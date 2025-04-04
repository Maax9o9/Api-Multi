package infrestructure

import (
    "Multi/src/hallDoor/application"
    "Multi/src/hallDoor/application/repositorys"
    "Multi/src/hallDoor/application/services"
    "Multi/src/hallDoor/infrestructure/adapters"
    "log"
)

func InitHallDoor() (*service.AlertHallDoorService, *service.ReceiveHallDoorService, *service.UpdateHallDoorService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "hall.door",    // Queue
        "sensor.hallDoor", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    hallDoorRepo := NewPostgres()

    getAllHallDoorUseCase := application.NewReceiveHallDoorUseCase(hallDoorRepo)
    getHallDoorByIDUseCase := application.NewReceiveHallDoorByIDUseCase(hallDoorRepo)
    alertHallDoorUseCase := application.NewAlertHallDoorUseCase(hallDoorRepo, rabbitRepo)
    updateHallDoorUseCase := application.NewUpdateHallDoorUseCase(hallDoorRepo)

    alertHallDoorService := service.NewAlertHallDoorService(alertHallDoorUseCase)
    receiveHallDoorService := service.NewReceiveHallDoorService(getAllHallDoorUseCase, getHallDoorByIDUseCase)
    updateHallDoorService := service.NewUpdateHallDoorService(updateHallDoorUseCase)

    return alertHallDoorService, receiveHallDoorService, updateHallDoorService, rabbitMQ
}