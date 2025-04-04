package controllers

import (
    "bytes"
    "encoding/json"
    service "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertGasController struct {
    gasService *service.AlertGasService
}

func NewAlertGasController(gasService *service.AlertGasService) *AlertGasController {
    return &AlertGasController{
        gasService: gasService,
    }
}

func (agc *AlertGasController) CreateGasData(ctx *gin.Context) {
    var request entities.GasSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := agc.gasService.CreateGasData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create gas data",
            "details": err.Error(),
        })
        return
    }

    go func() {
        url := "http://localhost:7070/gas"
        payload := map[string]interface{}{
            "id":    request.ID,
            "created_at":      request.CreatedAt,
            "status":    request.Status,
            "gas_level": request.GasLevel,
            
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending gas data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send gas data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("Gas data sent successfully!")
        } else {
            println("Failed to send gas data. Status code:", resp.StatusCode)
        }
    }()

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Gas data created successfully",
    })
}