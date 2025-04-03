package controllers

import (
    service "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertGasController struct {
    gasService *service.AlertGasService
}

func NewAlertGasController(gasService *service.AlertGasService) *AlertGasController {
    return &AlertGasController{
        gasService: gasService,
    }
}

func (agc *AlertGasController) CreateGasData(ctx *gin.Context) {
    var request entities.GasSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := agc.gasService.CreateGasData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create gas data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Gas data created successfully",
    })
}