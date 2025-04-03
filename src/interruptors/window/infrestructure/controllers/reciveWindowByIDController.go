package controllers

import (
    "Multi/src/interruptors/window/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveWindowByIDController struct {
    windowService *service.ReceiveWindowService
}

func NewReceiveWindowByIDController(windowService *service.ReceiveWindowService) *ReceiveWindowByIDController {
    return &ReceiveWindowByIDController{
        windowService: windowService,
    }
}

func (rwc *ReceiveWindowByIDController) GetWindowDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := rwc.windowService.GetWindowDataByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
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