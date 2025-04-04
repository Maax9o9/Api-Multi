package controllers

import (
    "Multi/src/hallDoor/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveHallDoorController struct {
    hallDoorService *service.ReceiveHallDoorService
}

func NewReceiveHallDoorController(hallDoorService *service.ReceiveHallDoorService) *ReceiveHallDoorController {
    return &ReceiveHallDoorController{
        hallDoorService: hallDoorService,
    }
}

func (rhdc *ReceiveHallDoorController) GetAllHallDoors(c *gin.Context) {
    data, err := rhdc.hallDoorService.GetAllHallDoors()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al obtener los datos de HallDoor",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}

func (rhdc *ReceiveHallDoorController) GetLatestHallDoorData(c *gin.Context) {
    data, err := rhdc.hallDoorService.GetLatestHallDoorData()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "No hay datos recientes de HallDoor",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}