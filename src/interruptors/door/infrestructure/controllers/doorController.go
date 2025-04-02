package controllers

import (
    "Multi/src/interruptors/door/application/services"
    "Multi/src/interruptors/door/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DoorController struct {
    doorService *service.DoorService
}

func NewDoorController(doorService *service.DoorService) *DoorController {
    return &DoorController{
        doorService: doorService,
    }
}

func (dc *DoorController) GetAllDoorData(ctx *gin.Context) {
    data, err := dc.doorService.GetAllDoorData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to get door data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Door data retrieved successfully",
        "data":    data,
    })
}

func (dc *DoorController) GetLatestDoorData(ctx *gin.Context) {
    data, err := dc.doorService.GetLatestDoorData()
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "No recent door data available",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Latest door data retrieved successfully",
        "data":    data,
    })
}

func (dc *DoorController) CreateDoorData(ctx *gin.Context) {
    var request entities.DoorData
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := dc.doorService.CreateDoorData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create door data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Door data created successfully",
    })
}

func (dc *DoorController) GetDoorDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := dc.doorService.GetDoorDataByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "Failed to get door data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Door data retrieved successfully",
        "data":    data,
    })
}