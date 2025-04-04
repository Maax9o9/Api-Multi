package infrestructure

import (
	"Multi/src/weather/application"
	"Multi/src/weather/application/repositorys"
	service "Multi/src/weather/application/services"
	"Multi/src/weather/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func InitWeather() (*service.AlertWeatherService, *service.ReceiveWeatherService, *adapters.RabbitConsumer) {
	rabbitMQ, err := adapters.NewRabbitConsumer(
		"amqp://uriel:eduardo117@3.228.81.226:5672/",
		"amq.topic",    // Exchange
		"temp.damp",    // Queue
		"sensor.dht11", // Routing Key
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	go func() {
		log.Println("Iniciando consumo de mensajes de RabbitMQ...")
		err := rabbitMQ.ConsumeMessages(func(body []byte) {
			log.Printf("Mensaje recibido de RabbitMQ: %s\n", string(body))

			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error al parsear mensaje JSON: %v", err)
				return
			}

			var temperatura, humedad float64

			if temp, ok := data["temperatura"].(float64); ok {
				temperatura = temp
			} else {
				log.Println("No se encontró valor válido para temperatura")
			}

			if hum, ok := data["humedad"].(float64); ok {
				humedad = hum
			} else {
				log.Println("No se encontró valor válido para humedad")
			}

			go sendWeatherDataToWebSocket(temperatura, humedad)
		})

		if err != nil {
			log.Printf("Error al consumir mensajes: %v", err)
		}
	}()

	rabbitRepo := repositorys.NewRabbitRepository(rabbitMQ)
	weatherRepo := NewPostgres()

	getAllWeatherUseCase := application.NewReceiveWeatherUseCase(weatherRepo)
	getWeatherByIDUseCase := application.NewReceiveWeatherByIDUseCase(weatherRepo)

	alertWeatherUseCase := application.NewAlertWeatherUseCase(weatherRepo, rabbitRepo)
	alertWeatherService := service.NewAlertWeatherService(alertWeatherUseCase)
	receiveWeatherService := service.NewReceiveWeatherService(getAllWeatherUseCase, getWeatherByIDUseCase)

	return alertWeatherService, receiveWeatherService, rabbitMQ
}

func sendWeatherDataToWebSocket(temperatura, humedad float64) {
	url := "http://35.171.234.157:7070/weather"

	payload := map[string]interface{}{
		"weather_id": 0,
		"date":       time.Now(),
		"heat":       temperatura,
		"damp":       humedad,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al crear JSON para WebSocket: %v", err)
		return
	}

	log.Println("Enviando datos meteorológicos al WebSocket:")
	log.Println("URL:", url)
	log.Println("Payload:", string(jsonData))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al enviar notificación: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Datos meteorológicos enviados correctamente")
	} else {
		log.Printf("Error al enviar datos. Código de estado: %d", resp.StatusCode)

		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr == nil && len(bodyBytes) > 0 {
			log.Printf("Respuesta del servidor: %s", string(bodyBytes))
		}
	}
}
