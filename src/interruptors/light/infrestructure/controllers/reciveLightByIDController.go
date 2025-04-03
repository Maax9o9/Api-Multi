package controllers

import (
    "Multi/src/interruptors/light/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveLightByIDController struct {
    lightService *service.ReceiveLightService
}

func NewReceiveLightByIDController(lightService *service.ReceiveLightService) *ReceiveLightByIDController {
    return &ReceiveLightByIDController{
        lightService: lightService,
    }
}

func (rlc *ReceiveLightByIDController) GetLightDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := rlc.lightService.GetLightDataByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
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