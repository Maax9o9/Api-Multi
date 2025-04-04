package application

import (
	"Multi/src/interruptors/lightOutside/application/repositorys"
	"Multi/src/interruptors/lightOutside/domain"
	"Multi/src/interruptors/lightOutside/domain/entities"
	"log"
	"strconv"
)

type AlertLightUseCase struct {
	lightRepo  domain.LightOutsideRepository
	rabbitRepo *repositorys.RabbitRepository // Corregido: Usa el tipo concreto, no el de dominio
}

func NewAlertLightUseCase(lightRepo domain.LightOutsideRepository, rabbitRepo *repositorys.RabbitRepository) *AlertLightUseCase {
	return &AlertLightUseCase{
		lightRepo:  lightRepo,
		rabbitRepo: rabbitRepo,
	}
}

// CreateLightData crea un nuevo registro de iluminación y envía el comando a través de RabbitMQ
func (uc *AlertLightUseCase) CreateLightData(lightData *entities.LightOutsideData) error {
	// Guardar en la base de datos
	err := uc.lightRepo.Create(lightData)
	if err != nil {
		return err
	}

	// Enviar comando al sistema de iluminación a través de RabbitMQ
	statusStr := strconv.Itoa(lightData.Status)
	// Corregido: Llamando al método SendLightOutsideCommand
	err = uc.rabbitRepo.SendLightOutsideCommand([]byte(statusStr))
	if err != nil {
		log.Printf("Error al enviar comando a RabbitMQ, pero continuando: %v", err)
	}

	return nil
}
