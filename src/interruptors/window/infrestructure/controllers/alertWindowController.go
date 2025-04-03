package controllers

import (
    "Multi/src/interruptors/window/application/services"
    "Multi/src/interruptors/window/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertWindowController struct {
    windowService *service.AlertWindowService
}

func NewAlertWindowController(windowService *service.AlertWindowService) *AlertWindowController {
    return &AlertWindowController{
        windowService: windowService,
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

    err := awc.windowService.CreateWindowData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create window data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Window data created successfully",
    })
}