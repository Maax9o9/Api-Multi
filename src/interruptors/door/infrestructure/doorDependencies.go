package infrestructure

import (
    "Multi/src/interruptors/door/application"
    "Multi/src/interruptors/door/application/repositorys"
    "Multi/src/interruptors/door/application/services"
    "Multi/src/interruptors/door/infrestructure/adapters"
    "log"
)

func InitDoor() (*service.AlertDoorService, *service.ReceiveDoorService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "puerta",       // Queue
        "sensor.puerta", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    doorRepo := NewPostgres()

    getAllDoorUseCase := application.NewReceiveDoorUseCase(doorRepo)
    getDoorByIDUseCase := application.NewGetDoorByIDUseCase(doorRepo)

    alertDoorUseCase := application.NewAlertDoorUseCase(doorRepo, rabbitRepo)
    alertDoorService := service.NewAlertDoorService(alertDoorUseCase)
    receiveDoorService := service.NewReceiveDoorService(getAllDoorUseCase, getDoorByIDUseCase)

    return alertDoorService, receiveDoorService, rabbitMQ
}