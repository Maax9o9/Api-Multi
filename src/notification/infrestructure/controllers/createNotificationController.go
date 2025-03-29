package controllers

import (
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
        HouseID          int    `json:"house_id"`
        Message          string `json:"message"`
        TypeNotification string `json:"type_notification"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    notification, err := c.useCase.CreateNotification(request.HouseID, request.Message, request.TypeNotification)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, notification)
}