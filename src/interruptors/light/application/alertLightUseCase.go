package application

import (
	"Multi/src/interruptors/light/application/repositorys"
	"Multi/src/interruptors/light/domain"
	"Multi/src/interruptors/light/domain/entities"
	"log"
	"strconv"
)

type AlertLightUseCase struct {
	lightRepo  domain.LightRepository
	rabbitRepo *repositorys.RabbitRepository
}

func NewAlertLightUseCase(lightRepo domain.LightRepository, rabbitRepo *repositorys.RabbitRepository) *AlertLightUseCase {
	return &AlertLightUseCase{
		lightRepo:  lightRepo,
		rabbitRepo: rabbitRepo,
	}
}

// CreateLightData crea un nuevo registro de iluminación y envía el comando a través de RabbitMQ
func (uc *AlertLightUseCase) CreateLightData(lightData *entities.LightData) error {
	// Guardar en la base de datos
	err := uc.lightRepo.Create(lightData)
	if err != nil {
		return err
	}

	// Enviar comando al sistema de iluminación a través de RabbitMQ
	statusStr := strconv.Itoa(lightData.Status)
	err = uc.rabbitRepo.SendLightCommand([]byte(statusStr))
	if err != nil {
		log.Printf("Error al enviar comando a RabbitMQ, pero continuando: %v", err)
	}

	return nil
}
