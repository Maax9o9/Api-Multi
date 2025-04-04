package application

import (
	"Multi/src/interruptors/door/domain"
	"Multi/src/interruptors/door/domain/entities"
	"log"
	"strconv"
)

type AlertDoorUseCase struct {
	doorRepo      domain.DoorRepository
	messagingRepo domain.MessagingRepository // Usar la interfaz en lugar del tipo concreto
}

func NewAlertDoorUseCase(doorRepo domain.DoorRepository, messagingRepo domain.MessagingRepository) *AlertDoorUseCase {
	return &AlertDoorUseCase{
		doorRepo:      doorRepo,
		messagingRepo: messagingRepo,
	}
}

func (uc *AlertDoorUseCase) Create(doorData *entities.DoorData) error {
	// Guardar en la base de datos
	err := uc.doorRepo.Create(doorData)
	if err != nil {
		return err
	}

	// Enviar comando a la puerta
	statusStr := strconv.Itoa(doorData.Status)
	err = uc.messagingRepo.SendDoorCommand([]byte(statusStr))
	if err != nil {
		log.Printf("Error al enviar comando a la puerta: %v", err)
		// Decide si quieres manejar este error o simplemente loggearlo
	}

	return nil
}
