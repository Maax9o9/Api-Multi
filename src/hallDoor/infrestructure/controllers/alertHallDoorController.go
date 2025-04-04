package controllers

import (
    "bytes"
    "encoding/json"
    "Multi/src/hallDoor/application/services"
    "Multi/src/hallDoor/domain/entities"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertHallDoorController struct {
    hallDoorService *service.AlertHallDoorService
}

func NewAlertHallDoorController(hallDoorService *service.AlertHallDoorService) *AlertHallDoorController {
    return &AlertHallDoorController{
        hallDoorService: hallDoorService,
    }
}

func (ahdc *AlertHallDoorController) ProcessHallDoorData(ctx *gin.Context) {
    var hallDoorData entities.HallDoor

    if err := ctx.ShouldBindJSON(&hallDoorData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request format",
            "details": err.Error(),
        })
        return
    }

    err := ahdc.hallDoorService.CreateHallDoor(&hallDoorData)
    if err != nil {
        log.Printf("Error al guardar los datos de HallDoor: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to save HallDoor data",
            "details": err.Error(),
        })
        return
    }

    go func() {
        url := "http://localhost:7070/hall-door"
        payload := map[string]interface{}{
            "id": hallDoorData.ID,
            "date":         hallDoorData.Date,
            "status":       hallDoorData.Status,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending HallDoor data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send HallDoor data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("HallDoor data sent successfully!")
        } else {
            println("Failed to send HallDoor data. Status code:", resp.StatusCode)
        }
    }()

    log.Printf("Datos de HallDoor procesados y guardados: %+v", hallDoorData)
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "HallDoor data processed and saved successfully",
    })
}