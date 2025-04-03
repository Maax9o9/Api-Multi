package controllers

import (
    "Multi/src/interruptors/door/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveDoorController struct {
    doorService *service.ReceiveDoorService
}

func NewReceiveDoorController(doorService *service.ReceiveDoorService) *ReceiveDoorController {
    return &ReceiveDoorController{
        doorService: doorService,
    }
}

func (rdc *ReceiveDoorController) GetAllDoorData(ctx *gin.Context) {
    data, err := rdc.doorService.GetAllDoorData()
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

func (rdc *ReceiveDoorController) GetDoorDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := rdc.doorService.GetDoorDataByID(id)
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