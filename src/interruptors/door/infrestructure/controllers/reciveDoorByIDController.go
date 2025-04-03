package controllers

import (
    "Multi/src/interruptors/door/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type GetDoorByIDController struct {
    doorService *service.ReceiveDoorService
}

func NewGetDoorByIDController(doorService *service.ReceiveDoorService) *GetDoorByIDController {
    return &GetDoorByIDController{
        doorService: doorService,
    }
}

func (gdc *GetDoorByIDController) GetDoorDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    data, err := gdc.doorService.GetDoorDataByID(id)
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