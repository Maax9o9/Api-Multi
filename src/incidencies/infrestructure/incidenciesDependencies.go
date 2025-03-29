package infrestructure

import (
    "Multi/src/incidencies/application"
    "Multi/src/incidencies/application/repositorys"
    "Multi/src/incidencies/application/services"
    "Multi/src/incidencies/infrestructure/adapters"
    "log"
)

func InitIncidencies() (
    *application.GetIncidenciesUseCase,
    *application.IncrementIncidencyUseCase,
    *adapters.RabbitGasAdapter,
    *adapters.RabbitMovementAdapter,
    *adapters.RabbitDoorAdapter,
    *adapters.RabbitWindowAdapter,
    *service.IncidenciesService,
    *adapters.RabbitAdapter,
) {
    rabbitGas, err := adapters.NewRabbitGasAdapter("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ for gas: %v", err)
    }

    rabbitMovement, err := adapters.NewRabbitMovementAdapter("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ for movement: %v", err)
    }

    rabbitDoor, err := adapters.NewRabbitDoorAdapter("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ for door: %v", err)
    }

    rabbitWindow, err := adapters.NewRabbitWindowAdapter("amqp://user:password@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ for window: %v", err)
    }

    rabbitMQ, err := adapters.NewRabbitAdapter("amqp://user:password@localhost:5672/", "IncidenciesQueue")
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    incidenciesRepo := NewPostgres()

    getIncidenciesUseCase := application.NewGetIncidenciesUseCase(incidenciesRepo, rabbitRepo)
    incrementIncidencyUseCase := application.NewIncrementIncidencyUseCase(incidenciesRepo, rabbitRepo)

    incidenciesService := service.NewIncidenciesService(getIncidenciesUseCase, incrementIncidencyUseCase)

    return getIncidenciesUseCase, incrementIncidencyUseCase, rabbitGas, rabbitMovement, rabbitDoor, rabbitWindow, incidenciesService, rabbitMQ
}