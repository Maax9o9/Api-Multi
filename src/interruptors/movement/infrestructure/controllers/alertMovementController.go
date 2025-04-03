package controllers

import (
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

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Movement data created successfully",
    })
}