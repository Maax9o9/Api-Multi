package controllers

import (
    "Multi/src/interruptors/door/application"
    "Multi/src/interruptors/door/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type DoorController struct {
    useCase *application.DoorUseCase
}

func NewDoorController(useCase *application.DoorUseCase) *DoorController {
    return &DoorController{
        useCase: useCase,
    }
}

func (c *DoorController) SendDoorCommand(ctx *gin.Context) {
    var command entities.DoorCommand
    if err := ctx.ShouldBindJSON(&command); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    if command.State != 0 && command.State != 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "State must be 0 or 1"})
        return
    }

    err := c.useCase.SendDoorCommand(command)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send door command", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Door command sent successfully"})
}