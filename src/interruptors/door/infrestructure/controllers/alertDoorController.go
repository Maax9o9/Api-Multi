package controllers

import (
    "Multi/src/interruptors/door/application/services"
    "Multi/src/interruptors/door/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertDoorController struct {
    doorService *service.AlertDoorService
}

func NewAlertDoorController(doorService *service.AlertDoorService) *AlertDoorController {
    return &AlertDoorController{
        doorService: doorService,
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

    err := adc.doorService.CreateDoorData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create door data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Door data created successfully",
    })
}