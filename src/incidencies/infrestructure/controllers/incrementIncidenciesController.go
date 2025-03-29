package controllers

import (
    "Multi/src/incidencies/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type IncrementIncidenciesController struct {
    useCase *application.IncrementIncidencyUseCase
}

func NewIncrementIncidenciesController(useCase *application.IncrementIncidencyUseCase) *IncrementIncidenciesController {
    return &IncrementIncidenciesController{
        useCase: useCase,
    }
}

func (c *IncrementIncidenciesController) IncrementIncidency(ctx *gin.Context) {
    typeNotification := ctx.Param("type")
    _, err := c.useCase.IncrementIncidency(typeNotification)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al incrementar la incidencia", "details": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Incidencia incrementada exitosamente"})
}