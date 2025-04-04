package controllers

import (
    "Multi/src/hallDoor/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveHallDoorByIDController struct {
    hallDoorService *service.ReceiveHallDoorService
}

func NewReceiveHallDoorByIDController(hallDoorService *service.ReceiveHallDoorService) *ReceiveHallDoorByIDController {
    return &ReceiveHallDoorByIDController{
        hallDoorService: hallDoorService,
    }
}

func (rhdc *ReceiveHallDoorByIDController) GetHallDoorByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
        })
        return
    }

    data, err := rhdc.hallDoorService.GetHallDoorByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "Error al obtener los datos de HallDoor",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, data)
}