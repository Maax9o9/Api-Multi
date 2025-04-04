package repositorys

import (
	"Multi/src/interruptors/lightOutside/infrestructure/adapters"
	"log"
)

type RabbitRepository struct {
	publisher *adapters.RabbitPublisher
}

// NewRabbitRepository crea un nuevo repositorio que utiliza RabbitMQ para enviar comandos
func NewRabbitRepository(publisher *adapters.RabbitPublisher) *RabbitRepository {
	return &RabbitRepository{publisher: publisher}
}

// SendLightOutsideCommand envía un comando al sistema de iluminación exterior a través de RabbitMQ
func (repo *RabbitRepository) SendLightOutsideCommand(command []byte) error {
	// La routing key específica para comandos de iluminación exterior
	routingKey := "light.out"

	// Publicar el comando
	err := repo.publisher.PublishMessage(routingKey, command)
	if err != nil {
		log.Printf("Error al enviar comando de iluminación exterior: %v", err)
		return err
	}

	log.Printf("Comando de iluminación exterior enviado correctamente: %s", string(command))
	return nil
}
