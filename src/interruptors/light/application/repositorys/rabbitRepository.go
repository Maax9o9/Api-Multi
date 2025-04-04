package repositorys

import (
	"Multi/src/interruptors/light/infrestructure/adapters"
	"log"
)

type RabbitRepository struct {
	publisher *adapters.RabbitPublisher
}

// NewRabbitRepository crea un nuevo repositorio que utiliza RabbitMQ para enviar comandos
func NewRabbitRepository(publisher *adapters.RabbitPublisher) *RabbitRepository {
	return &RabbitRepository{publisher: publisher}
}

// SendLightCommand envía un comando al sistema de iluminación a través de RabbitMQ
func (repo *RabbitRepository) SendLightCommand(command []byte) error {
	// La routing key específica para comandos de iluminación
	routingKey := "light.on"

	// Publicar el comando
	err := repo.publisher.PublishMessage(routingKey, command)
	if err != nil {
		log.Printf("Error al enviar comando de iluminación: %v", err)
		return err
	}

	log.Printf("Comando de iluminación enviado correctamente: %s", string(command))
	return nil
}

// Mantenemos el método ProcessLightCommands para compatibilidad con código existente
func (repo *RabbitRepository) ProcessLightCommands(processMessage func(body []byte)) error {
	log.Printf("Esta función está obsoleta. Usa SendLightCommand en su lugar.")
	return nil
}
