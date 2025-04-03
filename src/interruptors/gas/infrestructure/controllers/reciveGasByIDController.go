package controllers

import (
    service "Multi/src/interruptors/gas/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveGasByIDController struct {
    gasService *service.ReceiveGasService
}

func NewReceiveGasByIDController(gasService *service.ReceiveGasService) *ReceiveGasByIDController {
    return &ReceiveGasByIDController{
        gasService: gasService,
    }
}

func (rgc *ReceiveGasByIDController) GetGasDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := rgc.gasService.GetGasDataByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
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
