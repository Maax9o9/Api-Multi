package infrestructure

import (
    "Multi/src/interruptors/door/application"
    "Multi/src/interruptors/door/application/repositorys"
    "Multi/src/interruptors/door/application/services"
    "Multi/src/interruptors/door/infrestructure/adapters"
    "log"
)

func InitDoor() (*service.DoorService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",                         
        "puerta",                            
        "sensor.puerta",                       
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    doorRepo := NewPostgres()

    doorUseCase := application.NewDoorUseCase(doorRepo, rabbitRepo)

    doorService := service.NewDoorService(doorUseCase)

    return doorService, rabbitMQ
}