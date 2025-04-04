package application

import (
	"Multi/src/interruptors/window/application/repositorys"
	"Multi/src/interruptors/window/domain"
	"Multi/src/interruptors/window/domain/entities"
	"log"
	"strconv"
)

type AlertWindowUseCase struct {
	windowRepo domain.WindowRepository
	rabbitRepo *repositorys.RabbitRepository
}

func NewAlertWindowUseCase(windowRepo domain.WindowRepository, rabbitRepo *repositorys.RabbitRepository) *AlertWindowUseCase {
	return &AlertWindowUseCase{
		windowRepo: windowRepo,
		rabbitRepo: rabbitRepo,
	}
}

// Create implementa el método requerido por la interfaz (asegúrate de que coincida con tu interfaz)
func (uc *AlertWindowUseCase) CreateWindowData(windowData *entities.WindowSensor) error {
	// Guardar en la base de datos
	err := uc.windowRepo.Create(windowData)
	if err != nil {
		return err
	}

	// Enviar comando a la ventana a través de RabbitMQ
	statusStr := strconv.Itoa(windowData.Status)
	err = uc.rabbitRepo.SendWindowCommand([]byte(statusStr))
	if err != nil {
		// Aquí puedes elegir si quieres manejar el error o solo loggearlo
		log.Printf("Error al enviar comando a RabbitMQ, pero continuando: %v", err)
	}

	return nil
}
