package controllers

import (
    "Multi/src/interruptors/light/application"
    "Multi/src/interruptors/light/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type LightController struct {
    useCase *application.LightUseCase
}

func NewLightController(useCase *application.LightUseCase) *LightController {
    return &LightController{
        useCase: useCase,
    }
}

func (c *LightController) SendLightCommand(ctx *gin.Context) {
    var command entities.LightCommand
    if err := ctx.ShouldBindJSON(&command); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    if command.State != 0 && command.State != 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "State must be 0 or 1"})
        return
    }

    err := c.useCase.SendLightCommand(command)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send light command", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Light command sent successfully"})
}