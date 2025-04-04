package controllers

import (
    "Multi/src/hallWindow/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveHallWindowByIDController struct {
    hallWindowService *service.ReceiveHallWindowService
}

func NewReceiveHallWindowByIDController(hallWindowService *service.ReceiveHallWindowService) *ReceiveHallWindowByIDController {
    return &ReceiveHallWindowByIDController{
        hallWindowService: hallWindowService,
    }
}

func (rhwc *ReceiveHallWindowByIDController) GetHallWindowByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
        })
        return
    }

    data, err := rhwc.hallWindowService.GetHallWindowByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "Error al obtener los datos de HallWindow",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, data)
}