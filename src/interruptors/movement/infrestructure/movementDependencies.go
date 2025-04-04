package infrestructure

import (
	"Multi/src/interruptors/movement/application"
	"Multi/src/interruptors/movement/application/repositorys"
	service "Multi/src/interruptors/movement/application/services"
	"Multi/src/interruptors/movement/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func InitMovement() (*service.AlertMovementService, *service.ReceiveMovementService, *adapters.RabbitConsumer) {
	rabbitMQ, err := adapters.NewRabbitConsumer(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic",  // Exchange
		"movimiento", // Queue
		"sensor.pir", // Routing Key
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	go func() {
		log.Println("Iniciando consumo de mensajes de movimiento desde RabbitMQ...")
		err := rabbitMQ.ConsumeMessages(func(body []byte) {
			// Registrar el mensaje recibido
			log.Printf("Mensaje de movimiento recibido: %s\n", string(body))

			// Parsear el mensaje JSON
			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error al parsear mensaje JSON del sensor de movimiento: %v", err)
				return
			}

			// Para un sensor PIR, generalmente cualquier mensaje indica detección de movimiento
			// Por lo tanto, establecemos status = 1 (movimiento detectado)
			status := 1

			// Si el mensaje contiene un campo específico que indica el estado, podemos usarlo
			if message, ok := data["message"].(string); ok {
				log.Printf("Mensaje del sensor: %s", message)
				// Si el mensaje indica "No hay movimiento", podemos cambiar el estado
				if message == "No hay movimiento" {
					status = 0
				}
			}

			// Enviar los datos al WebSocket
			go sendMovementDataToWebSocket(0, status) // ID 0 o incremental si lo prefieres
		})

		if err != nil {
			log.Printf("Error al consumir mensajes de movimiento: %v", err)
		}
	}()
	rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
	movementRepo := NewPostgres()

	getAllMovementUseCase := application.NewReceiveMovementUseCase(movementRepo)
	getMovementByIDUseCase := application.NewReceiveMovementByIDUseCase(movementRepo)

	alertMovementUseCase := application.NewAlertMovementUseCase(movementRepo, rabbitRepo)
	alertMovementService := service.NewAlertMovementService(alertMovementUseCase)
	receiveMovementService := service.NewReceiveMovementService(getAllMovementUseCase, getMovementByIDUseCase)

	return alertMovementService, receiveMovementService, rabbitMQ
}

func sendMovementDataToWebSocket(id, status int) {
	url := "http://35.171.234.157:7070/motion" // Ajusta según la URL de tu endpoint

	// Crear el payload según la estructura esperada
	payload := map[string]interface{}{
		"id":         id,
		"created_at": time.Now(),
		"status":     status,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al crear JSON para WebSocket de movimiento: %v", err)
		return
	}

	log.Println("Enviando datos de movimiento al WebSocket:")
	log.Println("URL:", url)
	log.Println("Payload:", string(jsonData))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al enviar datos de movimiento: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Datos de movimiento enviados correctamente")
	} else {
		log.Printf("Error al enviar datos de movimiento. Código de estado: %d", resp.StatusCode)

		// Leer y mostrar el cuerpo de la respuesta para depuración
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr == nil && len(bodyBytes) > 0 {
			log.Printf("Respuesta del servidor: %s", string(bodyBytes))
		}
	}
}
