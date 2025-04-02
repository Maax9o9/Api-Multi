package controllers

import (
    "Multi/src/interruptors/light/application/services"
    "Multi/src/interruptors/light/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type LightController struct {
    lightService *service.LightService
}

func NewLightController(lightService *service.LightService) *LightController {
    return &LightController{
        lightService: lightService,
    }
}

func (lc *LightController) GetAllLightData(ctx *gin.Context) {
    data, err := lc.lightService.GetAllLightData()
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

func (lc *LightController) GetLatestLightData(ctx *gin.Context) {
    data, err := lc.lightService.GetLatestLightData()
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

func (lc *LightController) CreateLightData(ctx *gin.Context) {
    var request entities.LightData
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := lc.lightService.CreateLightData(&request)
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

func (lc *LightController) GetLightDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := lc.lightService.GetLightDataByID(id)
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