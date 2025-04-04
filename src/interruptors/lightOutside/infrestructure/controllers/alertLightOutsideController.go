package controllers

import (
    "bytes"
    "encoding/json"
    "Multi/src/interruptors/lightOutside/application/services"
    "Multi/src/interruptors/lightOutside/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertLightController struct {
    lightService *service.AlertLightService
}

func NewAlertLightController(lightService *service.AlertLightService) *AlertLightController {
    return &AlertLightController{
        lightService: lightService,
    }
}

func (alc *AlertLightController) CreateLightData(ctx *gin.Context) {
    var request entities.LightOutsideData
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := alc.lightService.CreateLightData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create light data",
            "details": err.Error(),
        })
        return
    }

    go func() {
        url := "http://54.160.249.225:7070/light-outside"
        payload := map[string]interface{}{
            "id": request.ID,
            "created_at": request.CreatedAt,
            "status": request.Status,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending light outside data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send light outside data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("Light outside data sent successfully!")
        } else {
            println("Failed to send light outside data. Status code:", resp.StatusCode)
        }
    }()

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Light outside data created successfully",
    })
}