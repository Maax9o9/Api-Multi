package controllers

import (
    "bytes"
    "encoding/json"
    "Multi/src/notification/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateNotificationController struct {
    useCase *application.CreateNotificationUseCase
}

func NewCreateNotificationController(useCase *application.CreateNotificationUseCase) *CreateNotificationController {
    return &CreateNotificationController{
        useCase: useCase,
    }
}

func (c *CreateNotificationController) CreateNotification(ctx *gin.Context) {
    var request struct {
        SensorID   int    `json:"sensor_id"`
        SensorType string `json:"sensor_type"`
        Message    string `json:"message"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    notification, err := c.useCase.CreateNotification(request.SensorID, request.SensorType, request.Message)
    if err != nil {
        if err.Error() == "invalid sensorType: must be one of 'GasSensor', 'MotionSensor', 'DoorSensor', 'WindowSensor', or 'LedControl'" {
            ctx.JSON(http.StatusBadRequest, gin.H{
                "error":   "Invalid sensor_type",
                "details": err.Error(),
            })
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{
                "error":   "Failed to create notification",
                "details": err.Error(),
            })
        }
        return
    }

    go func() {
        url := "http://localhost:7070/notify"
        payload := map[string]string{
            "message": request.Message,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending notification to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send notification:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("Notification sent successfully!")
        } else {
            println("Failed to send notification. Status code:", resp.StatusCode)
        }
    }()

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Notification created successfully",
        "data":    notification,
    })
}