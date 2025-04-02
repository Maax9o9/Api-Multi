package controllers

import (
    service "Multi/src/interruptors/movement/application/services"
    "Multi/src/interruptors/movement/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type MovementController struct {
    movementService *service.MovementService
}

func NewMovementController(movementService *service.MovementService) *MovementController {
    return &MovementController{
        movementService: movementService,
    }
}

func (mc *MovementController) GetAllMovementData(ctx *gin.Context) {
    data, err := mc.movementService.GetAllMovementData()
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

func (mc *MovementController) GetMovementDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := mc.movementService.GetMovementDataByID(id)
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

func (mc *MovementController) CreateMovementData(ctx *gin.Context) {
    var request entities.MotionSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := mc.movementService.CreateMovementData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create movement data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Movement data created successfully",
    })
}