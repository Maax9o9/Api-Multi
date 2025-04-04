package repositorys

import (
	"Multi/src/interruptors/window/infrestructure/adapters"
	"log"
)

type RabbitRepository struct {
	publisher *adapters.RabbitPublisher
}

// NewRabbitRepository crea un nuevo repositorio que utiliza RabbitMQ para enviar comandos
func NewRabbitRepository(publisher *adapters.RabbitPublisher) *RabbitRepository {
	return &RabbitRepository{publisher: publisher}
}

// SendDoorCommand envía un comando a la puerta a través de RabbitMQ
func (repo *RabbitRepository) SendWindowCommand(command []byte) error {
	// La routing key específica para comandos de la puerta
	routingKey := "sensor.hallWindow"

	// Publicar el comando
	err := repo.publisher.PublishMessage(routingKey, command)
	if err != nil {
		log.Printf("Error al enviar comando de puerta: %v", err)
		return err
	}

	log.Printf("Comando de puerta enviado correctamente: %s", string(command))
	return nil
}
