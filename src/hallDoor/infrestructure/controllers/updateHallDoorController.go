package controllers

import (
    "Multi/src/hallDoor/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type UpdateHallDoorController struct {
    hallDoorService *service.UpdateHallDoorService
}

func NewUpdateHallDoorController(hallDoorService *service.UpdateHallDoorService) *UpdateHallDoorController {
    return &UpdateHallDoorController{
        hallDoorService: hallDoorService,
    }
}

func (uhdc *UpdateHallDoorController) UpdateHallDoorStatus(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
        })
        return
    }

    var requestBody struct {
        Status int `json:"status"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Formato de solicitud inválido",
            "details": err.Error(),
        })
        return
    }

    err = uhdc.hallDoorService.UpdateStatus(id, requestBody.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al actualizar el estado de HallDoor",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Estado de HallDoor actualizado correctamente",
    })
}