package infrestructure

import (
    "Multi/src/weather/application"
    "Multi/src/weather/application/repositorys"
    "Multi/src/weather/application/services"
    "Multi/src/weather/infrestructure/adapters"
    "log"
)

func InitWeather() (*service.AlertWeatherService, *service.ReceiveWeatherService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",    // Exchange
        "temp.damp",    // Queue
        "sensor.dht11", // Routing Key
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    weatherRepo := NewPostgres()

    alertWeatherUseCase := application.NewAlertWeatherUseCase(weatherRepo, rabbitRepo)
    receiveWeatherUseCase := application.NewReceiveWeatherUseCase(weatherRepo)

    alertWeatherService := service.NewAlertWeatherService(alertWeatherUseCase)
    receiveWeatherService := service.NewReceiveWeatherService(receiveWeatherUseCase)

    return alertWeatherService, receiveWeatherService, rabbitMQ
}