package controllers

import (
    "Multi/src/interruptors/movement/application"
    "Multi/src/interruptors/movement/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type MovementController struct {
    useCase *application.MovementUseCase
}

func NewMovementController(useCase *application.MovementUseCase) *MovementController {
    return &MovementController{
        useCase: useCase,
    }
}

func (c *MovementController) SendMovementCommand(ctx *gin.Context) {
    var command entities.MovementCommand
    if err := ctx.ShouldBindJSON(&command); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    if command.State != 0 && command.State != 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "State must be 0 or 1"})
        return
    }

    err := c.useCase.SendMovementCommand(command)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send movement command", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Movement command sent successfully"})
}