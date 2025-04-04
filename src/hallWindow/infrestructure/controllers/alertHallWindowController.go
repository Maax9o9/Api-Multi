package controllers

import (
    "bytes"
    "encoding/json"
    "Multi/src/hallWindow/application/services"
    "Multi/src/hallWindow/domain/entities"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertHallWindowController struct {
    hallWindowService *service.AlertHallWindowService
}

func NewAlertHallWindowController(hallWindowService *service.AlertHallWindowService) *AlertHallWindowController {
    return &AlertHallWindowController{
        hallWindowService: hallWindowService,
    }
}

func (ahwc *AlertHallWindowController) ProcessHallWindowData(ctx *gin.Context) {
    var hallWindowData entities.HallWindow

    if err := ctx.ShouldBindJSON(&hallWindowData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request format",
            "details": err.Error(),
        })
        return
    }

    err := ahwc.hallWindowService.CreateHallWindow(&hallWindowData)
    if err != nil {
        log.Printf("Error al guardar los datos de HallWindow: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to save HallWindow data",
            "details": err.Error(),
        })
        return
    }

    go func() {
        url := "http://35.171.234.157:7070/hall-window" 
        payload := map[string]interface{}{
            "id": hallWindowData.ID,
            "date":           hallWindowData.Date,
            "status":         hallWindowData.Status,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending HallWindow data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send HallWindow data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("HallWindow data sent successfully!")
        } else {
            println("Failed to send HallWindow data. Status code:", resp.StatusCode)
        }
    }()

    log.Printf("Datos de HallWindow procesados y guardados: %+v", hallWindowData)
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "HallWindow data processed and saved successfully",
    })
}