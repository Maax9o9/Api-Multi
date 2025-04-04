package controllers

import (
    "bytes"
    "encoding/json"
    "Multi/src/interruptors/movement/application/services"
    "Multi/src/interruptors/movement/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertMovementController struct {
    movementService *service.AlertMovementService
}

func NewAlertMovementController(movementService *service.AlertMovementService) *AlertMovementController {
    return &AlertMovementController{
        movementService: movementService,
    }
}

func (amc *AlertMovementController) CreateMovementData(ctx *gin.Context) {
    var request entities.MotionSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := amc.movementService.CreateMovementData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create movement data",
            "details": err.Error(),
        })
        return
    }

    go func() {
        url := "http://54.160.249.225:7070/motion" 
        payload := map[string]interface{}{
            "id": request.ID,
            "created_at": request.CreatedAt,
            "status": request.Status,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending movement data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send movement data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("Movement data sent successfully!")
        } else {
            println("Failed to send movement data. Status code:", resp.StatusCode)
        }
    }()

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Movement data created successfully",
    })
}