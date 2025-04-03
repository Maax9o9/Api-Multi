package controllers

import (
    "Multi/src/interruptors/movement/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveMovementByIDController struct {
    movementService *service.ReceiveMovementService
}

func NewReceiveMovementByIDController(movementService *service.ReceiveMovementService) *ReceiveMovementByIDController {
    return &ReceiveMovementByIDController{
        movementService: movementService,
    }
}

func (rmc *ReceiveMovementByIDController) GetMovementDataByID(ctx *gin.Context) {
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