package infrestructure

import (
    "Multi/src/weather/application"
    "Multi/src/weather/application/repositorys"
    "Multi/src/weather/application/services"
    "Multi/src/weather/infrestructure/adapters"
    "log"
)

func InitWeather() (*service.WeatherService, *adapters.RabbitConsumer) {
    rabbitMQ, err := adapters.NewRabbitConsumer(
        "amqp://uriel:eduardo117@3.228.81.226:5672/",
        "amq.topic",
        "temp.damp",
        "sensor.dht11",
    )
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }

    rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
    weatherRepo := NewPostgres()

    weatherUseCase := application.NewWeatherUseCase(weatherRepo, rabbitRepo)

    weatherService := service.NewWeatherService(weatherUseCase)

    return weatherService, rabbitMQ
}