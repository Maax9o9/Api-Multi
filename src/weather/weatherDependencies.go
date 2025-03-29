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
        "amqp://max:123@44.213.165.25:5672/",
        "PizzasExchange",
        "PizzasQueue",
        "pizza.order",
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