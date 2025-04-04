package controllers

import (
	service "Multi/src/interruptors/door/application/services"
	"Multi/src/interruptors/door/domain/entities"
	"Multi/src/interruptors/door/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlertDoorController struct {
	doorService *service.AlertDoorService
	publisher   *adapters.MQTTPublisher
}

func NewAlertDoorController(doorService *service.AlertDoorService, publisher *adapters.MQTTPublisher) *AlertDoorController {
	return &AlertDoorController{
		doorService: doorService,
		publisher:   publisher,
	}
}

func (adc *AlertDoorController) CreateDoorData(ctx *gin.Context) {
	var request entities.DoorData
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

	// Guardar en la base de datos y enviar el comando (a través del servicio)
	err := adc.doorService.CreateDoorData(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create door data",
			"details": err.Error(),
		})
		return
	}

	// Enviar a WebSocket
	go func() {
		url := "http://35.171.234.157:7070/door"
		payload := map[string]interface{}{
			"id":         request.ID,
			"created_at": request.CreatedAt,
			"status":     request.Status,
		}
		jsonData, _ := json.Marshal(payload)

		log.Println("Sending door data to WebSocket:")
		log.Println("URL:", url)
		log.Println("Payload:", string(jsonData))

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Failed to send door data:", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Door data sent successfully!")
		} else {
			log.Println("Failed to send door data. Status code:", resp.StatusCode)
		}
	}()

	// También enviar directamente a MQTT para redundancia
	go func() {
		statusStr := strconv.Itoa(request.Status)
		err := adc.publisher.PublishMessage("home/door/command", []byte(statusStr))
		if err != nil {
			log.Printf("Error al enviar comando de puerta directamente a MQTT: %v", err)
		} else {
			log.Printf("Comando enviado directamente a la ESP32 vía MQTT: %s", statusStr)
		}
	}()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Door data created successfully",
		"data": gin.H{
			"id":     request.ID,
			"status": request.Status,
		},
	})
}
