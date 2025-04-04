package controllers

import (
	service "Multi/src/interruptors/light/application/services"
	"Multi/src/interruptors/light/domain/entities"
	"Multi/src/interruptors/light/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlertLightController struct {
	lightService *service.AlertLightService
	publisher    *adapters.RabbitPublisher
}

func NewAlertLightController(lightService *service.AlertLightService, publisher *adapters.RabbitPublisher) *AlertLightController {
	return &AlertLightController{
		lightService: lightService,
		publisher:    publisher,
	}
}

func (alc *AlertLightController) CreateLightData(ctx *gin.Context) {
	var request entities.LightData
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	// Validar que el status sea 0 o 1
	if request.Status != 0 && request.Status != 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "El status debe ser 0 o 1",
		})
		return
	}

	// Guardar en la base de datos a través del servicio
	err := alc.lightService.CreateLightData(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create light data",
			"details": err.Error(),
		})
		return
	}

	// Enviar a WebSocket
	go func() {
		url := "http://54.160.249.225:7070/light"
		payload := map[string]interface{}{
			"id":         request.ID,
			"created_at": request.CreatedAt,
			"status":     request.Status,
		}
		jsonData, _ := json.Marshal(payload)

		log.Println("Sending light data to WebSocket:")
		log.Println("URL:", url)
		log.Println("Payload:", string(jsonData))

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Failed to send light data:", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Light data sent successfully!")
		} else {
			log.Println("Failed to send light data. Status code:", resp.StatusCode)
		}
	}()

	// También enviar directamente a RabbitMQ para asegurar que la ESP32 reciba el comando
	go func() {
		statusStr := strconv.Itoa(request.Status)
		err := alc.publisher.PublishMessage("actuator.light", []byte(statusStr))
		if err != nil {
			log.Printf("Error al enviar comando de iluminación a través de RabbitMQ: %v", err)
		} else {
			log.Printf("Comando de iluminación enviado directamente a la ESP32: %s", statusStr)
		}
	}()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Light data created successfully",
		"data": gin.H{
			"id":     request.ID,
			"status": request.Status,
		},
	})
}
