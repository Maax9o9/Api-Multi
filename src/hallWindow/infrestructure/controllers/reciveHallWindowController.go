package controllers

import (
    "Multi/src/hallWindow/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveHallWindowController struct {
    hallWindowService *service.ReceiveHallWindowService
}

func NewReceiveHallWindowController(hallWindowService *service.ReceiveHallWindowService) *ReceiveHallWindowController {
    return &ReceiveHallWindowController{
        hallWindowService: hallWindowService,
    }
}

func (rhwc *ReceiveHallWindowController) GetAllHallWindows(c *gin.Context) {
    data, err := rhwc.hallWindowService.GetAllHallWindows()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al obtener los datos de HallWindow",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}

func (rhwc *ReceiveHallWindowController) GetLatestHallWindowData(c *gin.Context) {
    data, err := rhwc.hallWindowService.GetLatestHallWindowData()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "No hay datos recientes de HallWindow",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}