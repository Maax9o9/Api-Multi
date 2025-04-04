package repositorys

import (
	"Multi/src/interruptors/door/infrestructure/adapters"
	"log"
)

type MQTTRepository struct {
	publisher *adapters.MQTTPublisher
}

// NewMQTTRepository crea un nuevo repositorio que utiliza MQTT para enviar comandos
func NewMQTTRepository(publisher *adapters.MQTTPublisher) *MQTTRepository {
	return &MQTTRepository{publisher: publisher}
}

// SendDoorCommand envía un comando a la puerta a través de MQTT
func (repo *MQTTRepository) SendDoorCommand(command []byte) error {
	// El topic MQTT específico para comandos de la puerta
	topic := "sensor.hallDoor"

	// Publicar el comando
	err := repo.publisher.PublishMessage(topic, command)
	if err != nil {
		log.Printf("Error al enviar comando de puerta: %v", err)
		return err
	}

	log.Printf("Comando de puerta enviado correctamente por MQTT: %s", string(command))
	return nil
}
