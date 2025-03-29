package controllers

import (
    "Multi/src/incidencies/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ShowIncidenciesController struct {
    useCase *application.GetIncidenciesUseCase
}

func NewShowIncidenciesController(useCase *application.GetIncidenciesUseCase) *ShowIncidenciesController {
    return &ShowIncidenciesController{
        useCase: useCase,
    }
}

func (c *ShowIncidenciesController) GetAllIncidencies(ctx *gin.Context) {
    incidencies, err := c.useCase.GetAllIncidencies()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las incidencias", "details": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, incidencies)
}

func (c *ShowIncidenciesController) GetIncidencyByType(ctx *gin.Context) {
    typeNotification := ctx.Param("type")
    incidency, err := c.useCase.GetIncidencyByType(typeNotification)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Incidencia no encontrada", "details": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, incidency)
}