package controllers

import (
    service "Multi/src/interruptors/gas/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveGasController struct {
    gasService *service.ReceiveGasService
}

func NewReceiveGasController(gasService *service.ReceiveGasService) *ReceiveGasController {
    return &ReceiveGasController{
        gasService: gasService,
    }
}

func (rgc *ReceiveGasController) GetAllGasData(ctx *gin.Context) {
    data, err := rgc.gasService.GetAllGasData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to retrieve gas data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Gas data retrieved successfully",
        "data":    data,
    })
}