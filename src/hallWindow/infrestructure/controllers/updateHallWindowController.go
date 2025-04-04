package controllers

import (
    "Multi/src/hallWindow/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type UpdateHallWindowController struct {
    hallWindowService *service.UpdateHallWindowService
}

func NewUpdateHallWindowController(hallWindowService *service.UpdateHallWindowService) *UpdateHallWindowController {
    return &UpdateHallWindowController{
        hallWindowService: hallWindowService,
    }
}

func (uhwc *UpdateHallWindowController) UpdateHallWindowStatus(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
            "details": err.Error(),
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

    if requestBody.Status < 0 || requestBody.Status > 1 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "El estado debe ser 0 o 1",
        })
        return
    }

    err = uhwc.hallWindowService.UpdateStatus(id, requestBody.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al actualizar el estado de HallWindow",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Estado de HallWindow actualizado correctamente",
    })
}