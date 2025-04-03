package controllers

import (
    "Multi/src/interruptors/light/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveLightController struct {
    lightService *service.ReceiveLightService
}

func NewReceiveLightController(lightService *service.ReceiveLightService) *ReceiveLightController {
    return &ReceiveLightController{
        lightService: lightService,
    }
}

func (rlc *ReceiveLightController) GetAllLightData(ctx *gin.Context) {
    data, err := rlc.lightService.GetAllLightData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to retrieve light data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Light data retrieved successfully",
        "data":    data,
    })
}

func (rlc *ReceiveLightController) GetLatestLightData(ctx *gin.Context) {
    data, err := rlc.lightService.GetLatestLightData()
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "No recent light data available",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Latest light data retrieved successfully",
        "data":    data,
    })
}