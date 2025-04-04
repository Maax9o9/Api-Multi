package infrestructure

import (
	"Multi/src/interruptors/door/application"
	"Multi/src/interruptors/door/application/repositorys"
	service "Multi/src/interruptors/door/application/services"
	"Multi/src/interruptors/door/infrestructure/adapters"
	"log"
	"time"
)

func InitDoor() (*service.AlertDoorService, *service.ReceiveDoorService, *adapters.MQTTPublisher) {
	mqttPublisher, err := adapters.NewMQTTPublisher(
		"tcp://3.228.81.226:1883",             // URL de tu broker RabbitMQ con puerto MQTT
		"door-api-client"+time.Now().String(), // Client ID único
		"uriel",                               // Usuario de RabbitMQ
		"eduardo117",                          // Contraseña de RabbitMQ
	)
	if err != nil {
		log.Fatalf("Error al inicializar el publicador MQTT: %v", err)
	}

	mqttRepo := repositorys.NewMQTTRepository(mqttPublisher)
	doorRepo := NewPostgres()

	getAllDoorUseCase := application.NewReceiveDoorUseCase(doorRepo)
	getDoorByIDUseCase := application.NewGetDoorByIDUseCase(doorRepo)

	alertDoorUseCase := application.NewAlertDoorUseCase(doorRepo, mqttRepo)
	alertDoorService := service.NewAlertDoorService(alertDoorUseCase)
	receiveDoorService := service.NewReceiveDoorService(getAllDoorUseCase, getDoorByIDUseCase)

	return alertDoorService, receiveDoorService, mqttPublisher
}
