package controllers

import (
	service "Multi/src/interruptors/window/application/services"
	"Multi/src/interruptors/window/domain/entities"
	"Multi/src/interruptors/window/infrestructure/adapters"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlertWindowController struct {
	windowService *service.AlertWindowService
	publisher     *adapters.RabbitPublisher
}

func NewAlertWindowController(windowService *service.AlertWindowService, publisher *adapters.RabbitPublisher) *AlertWindowController {
	return &AlertWindowController{
		windowService: windowService,
		publisher:     publisher,
	}
}

func (awc *AlertWindowController) CreateWindowData(ctx *gin.Context) {
	var request entities.WindowSensor
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
	err := awc.windowService.CreateWindowData(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create window data",
			"details": err.Error(),
		})
		return
	}

	// Enviar a WebSocket
	go func() {
		url := "http://localhost:7070/window"
		payload := map[string]interface{}{
			"id":         request.ID,
			"created_at": request.CreatedAt,
			"status":     request.Status,
		}
		jsonData, _ := json.Marshal(payload)

		log.Println("Sending window data to WebSocket:")
		log.Println("URL:", url)
		log.Println("Payload:", string(jsonData))

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Failed to send window data:", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Window data sent successfully!")
		} else {
			log.Println("Failed to send window data. Status code:", resp.StatusCode)
		}
	}()

	// También enviar directamente a RabbitMQ para asegurar que la ESP32 reciba el comando
	go func() {
		statusStr := strconv.Itoa(request.Status)
		err := awc.publisher.PublishMessage("actuator.window", []byte(statusStr))
		if err != nil {
			log.Printf("Error al enviar comando a la ventana a través de RabbitMQ: %v", err)
		} else {
			log.Printf("Comando enviado directamente a la ESP32: %s", statusStr)
		}
	}()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Window data created successfully",
		"data": gin.H{
			"id":     request.ID,
			"status": request.Status,
		},
	})
}
