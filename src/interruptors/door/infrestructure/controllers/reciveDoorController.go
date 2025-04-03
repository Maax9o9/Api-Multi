package controllers

import (
    "Multi/src/interruptors/door/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetAllDoorController struct {
    doorService *service.ReceiveDoorService
}

func NewGetAllDoorController(doorService *service.ReceiveDoorService) *GetAllDoorController {
    return &GetAllDoorController{
        doorService: doorService,
    }
}

func (gdc *GetAllDoorController) GetAllDoorData(ctx *gin.Context) {
    data, err := gdc.doorService.GetAllDoorData()
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