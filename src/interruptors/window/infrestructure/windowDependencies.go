package infrestructure

import (
	"Multi/src/interruptors/window/application"
	"Multi/src/interruptors/window/application/repositorys"
	service "Multi/src/interruptors/window/application/services"
	"Multi/src/interruptors/window/infrestructure/adapters"
	"log"
)

func InitWindow() (*service.AlertWindowService, *service.ReceiveWindowService, *adapters.RabbitPublisher) {
	publisher, err := adapters.NewRabbitPublisher(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic", // Exchange
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}
	rabbitRepo := repositorys.NewRabbitRepository(publisher)
	windowRepo := NewPostgres()

	getAllWindowUseCase := application.NewReceiveWindowUseCase(windowRepo)
	getWindowByIDUseCase := application.NewReceiveWindowByIDUseCase(windowRepo)

	alertWindowUseCase := application.NewAlertWindowUseCase(windowRepo, rabbitRepo)
	alertWindowService := service.NewAlertWindowService(alertWindowUseCase)
	receiveWindowService := service.NewReceiveWindowService(getAllWindowUseCase, getWindowByIDUseCase)

	return alertWindowService, receiveWindowService, publisher
}
