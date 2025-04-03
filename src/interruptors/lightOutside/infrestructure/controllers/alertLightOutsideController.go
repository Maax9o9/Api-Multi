package controllers

import (
    "Multi/src/interruptors/lightOutside/application/services"
    "Multi/src/interruptors/lightOutside/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertLightController struct {
    lightService *service.AlertLightService
}

func NewAlertLightController(lightService *service.AlertLightService) *AlertLightController {
    return &AlertLightController{
        lightService: lightService,
    }
}

func (alc *AlertLightController) CreateLightData(ctx *gin.Context) {
    var request entities.LightOutsideData
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := alc.lightService.CreateLightData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create light data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Light data created successfully",
    })
}