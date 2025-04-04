package infrestructure

import (
	"Multi/src/interruptors/lightOutside/application"
	"Multi/src/interruptors/lightOutside/application/repositorys"
	service "Multi/src/interruptors/lightOutside/application/services"
	"Multi/src/interruptors/lightOutside/infrestructure/adapters"
	"log"
)

func InitLightOutside() (*service.AlertLightService, *service.ReceiveLightService, *adapters.RabbitPublisher) {
	publisher, err := adapters.NewRabbitPublisher(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic", // Exchange
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}

	rabbitRepoOutside := repositorys.NewRabbitRepository(publisher)
	lightOutsideRepo := NewPostgres()

	getAllLightOutsideUseCase := application.NewReceiveLightUseCase(lightOutsideRepo)
	getLightOutsideByIDUseCase := application.NewReceiveLightByIDUseCase(lightOutsideRepo)

	alertLightOutsideUseCase := application.NewAlertLightUseCase(lightOutsideRepo, rabbitRepoOutside)
	alertLightOutsideService := service.NewAlertLightService(alertLightOutsideUseCase)
	receiveLightOutsideService := service.NewReceiveLightService(getAllLightOutsideUseCase, getLightOutsideByIDUseCase)

	return alertLightOutsideService, receiveLightOutsideService, publisher
}
