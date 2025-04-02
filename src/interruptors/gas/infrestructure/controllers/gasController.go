package controllers

import (
    service "Multi/src/interruptors/gas/application/services"
    "Multi/src/interruptors/gas/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type GasController struct {
    gasService *service.GasService
}

func NewGasController(gasService *service.GasService) *GasController {
    return &GasController{
        gasService: gasService,
    }
}

func (gc *GasController) GetAllGasData(ctx *gin.Context) {
    data, err := gc.gasService.GetAllGasData()
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

func (gc *GasController) GetGasDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := gc.gasService.GetGasDataByID(id)
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

func (gc *GasController) CreateGasData(ctx *gin.Context) {
    var request entities.GasSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := gc.gasService.CreateGasData(&request)
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