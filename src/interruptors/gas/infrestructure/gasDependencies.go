package infrestructure

import (
	"Multi/src/interruptors/gas/application"
	"Multi/src/interruptors/gas/application/repositorys"
	service "Multi/src/interruptors/gas/application/services"
	"Multi/src/interruptors/gas/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func InitGas() (*service.AlertGasService, *service.ReceiveGasService, *adapters.RabbitConsumer) {
	rabbitMQ, err := adapters.NewRabbitConsumer(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic",  // Exchange
		"gas",        // Queue
		"sensor.mq2", // Routing Key
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	go func() {
		log.Println("Iniciando consumo de mensajes de gas desde RabbitMQ...")
		err := rabbitMQ.ConsumeMessages(func(body []byte) {
			// Registrar el mensaje recibido
			log.Printf("Mensaje de gas recibido: %s\n", string(body))

			// Parsear el mensaje JSON
			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error al parsear mensaje JSON del sensor de gas: %v", err)
				return
			}

			// Extraer el nivel de gas
			var gasLevel float64
			var status int

			// Dependiendo de cómo venga el dato en el mensaje, ajusta estos campos
			if level, ok := data["nivel_gas"].(float64); ok {
				gasLevel = level
			} else if level, ok := data["nivel"].(float64); ok {
				gasLevel = level
			} else {
				log.Println("No se encontró un valor válido para el nivel de gas")
				return
			}

			if s, ok := data["status"].(float64); ok {
				status = int(s)
			} else if gasLevel > 400 { // Umbral de ejemplo, ajusta según tus necesidades
				status = 1 // Peligroso
			} else {
				status = 0 // Normal
			}

			go sendGasDataToWebSocket(0, status, gasLevel) // ID 0 o autoincremental si lo prefieres
		})

		if err != nil {
			log.Printf("Error al consumir mensajes de gas: %v", err)
		}
	}()

	rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
	gasRepo := NewPostgres()

	getAllGasUseCase := application.NewReceiveGasUseCase(gasRepo)
	getGasByIDUseCase := application.NewReceiveGasByIDUseCase(gasRepo)

	alertGasUseCase := application.NewAlertGasUseCase(gasRepo, rabbitRepo)
	alertGasService := service.NewAlertGasService(alertGasUseCase)
	receiveGasService := service.NewReceiveGasService(getAllGasUseCase, getGasByIDUseCase)

	return alertGasService, receiveGasService, rabbitMQ
}
func sendGasDataToWebSocket(id, status int, gasLevel float64) {
	url := "http://35.171.234.157:7070/gas"

	// Crear el payload según la estructura esperada
	payload := map[string]interface{}{
		"id":         id,
		"created_at": time.Now(),
		"status":     status,
		"gas_level":  gasLevel,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al crear JSON para WebSocket de gas: %v", err)
		return
	}

	log.Println("Enviando datos de gas al WebSocket:")
	log.Println("URL:", url)
	log.Println("Payload:", string(jsonData))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al enviar datos de gas: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Datos de gas enviados correctamente")
	} else {
		log.Printf("Error al enviar datos de gas. Código de estado: %d", resp.StatusCode)

		// Leer y mostrar el cuerpo de la respuesta para ayudar en la depuración
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr == nil && len(bodyBytes) > 0 {
			log.Printf("Respuesta del servidor: %s", string(bodyBytes))
		}
	}
}
