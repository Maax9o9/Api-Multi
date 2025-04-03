package infrestructure

import (
    "Multi/src/interruptors/gas/application"
    "Multi/src/interruptors/gas/application/repositorys"
    "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/infrestructure/adapters"
    "log"
)

func InitGas() (*service.GasService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",                         
        "gas",                            
        "sensor.mq2",                        
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    gasRepo := NewPostgres()

    gasUseCase := application.NewGasUseCase(gasRepo, rabbitRepo)

    gasService := service.NewGasService(gasUseCase)

    return gasService, rabbitMQ
}