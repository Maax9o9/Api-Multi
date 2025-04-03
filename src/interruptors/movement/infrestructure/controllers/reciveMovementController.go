package controllers

import (
    "Multi/src/interruptors/movement/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveMovementController struct {
    movementService *service.ReceiveMovementService
}

func NewReceiveMovementController(movementService *service.ReceiveMovementService) *ReceiveMovementController {
    return &ReceiveMovementController{
        movementService: movementService,
    }
}

func (rmc *ReceiveMovementController) GetAllMovementData(ctx *gin.Context) {
    data, err := rmc.movementService.GetAllMovementData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to retrieve movement data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Movement data retrieved successfully",
        "data":    data,
    })
}

func (rmc *ReceiveMovementController) GetMovementDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := rmc.movementService.GetMovementDataByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "Failed to retrieve movement data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Movement data retrieved successfully",
        "data":    data,
    })
}