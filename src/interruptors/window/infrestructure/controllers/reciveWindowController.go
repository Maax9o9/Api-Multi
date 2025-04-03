package controllers

import (
    "Multi/src/interruptors/window/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveWindowController struct {
    windowService *service.ReceiveWindowService
}

func NewReceiveWindowController(windowService *service.ReceiveWindowService) *ReceiveWindowController {
    return &ReceiveWindowController{
        windowService: windowService,
    }
}

func (rwc *ReceiveWindowController) GetAllWindowData(ctx *gin.Context) {
    data, err := rwc.windowService.GetAllWindowData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to retrieve window data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Window data retrieved successfully",
        "data":    data,
    })
}