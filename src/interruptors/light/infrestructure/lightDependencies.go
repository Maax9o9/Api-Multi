package infrestructure

import (
	"Multi/src/interruptors/light/application"
	"Multi/src/interruptors/light/application/repositorys"
	service "Multi/src/interruptors/light/application/services"
	"Multi/src/interruptors/light/infrestructure/adapters"
	"log"
)

func InitLight() (*service.AlertLightService, *service.ReceiveLightService, *adapters.RabbitPublisher) {
	publisher, err := adapters.NewRabbitPublisher(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic", // Exchange
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}

	rabbitRepo := repositorys.NewRabbitRepository(publisher)
	lightRepo := NewPostgres()

	getAllLightUseCase := application.NewReceiveLightUseCase(lightRepo)
	getLightByIDUseCase := application.NewReceiveLightByIDUseCase(lightRepo)

	alertLightUseCase := application.NewAlertLightUseCase(lightRepo, rabbitRepo)
	alertLightService := service.NewAlertLightService(alertLightUseCase)
	receiveLightService := service.NewReceiveLightService(getAllLightUseCase, getLightByIDUseCase)

	return alertLightService, receiveLightService, publisher
}
