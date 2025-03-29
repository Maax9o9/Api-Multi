package controllers

import (
    "Multi/src/interruptors/window/application"
    "Multi/src/interruptors/window/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type WindowController struct {
    useCase *application.WindowUseCase
}

func NewWindowController(useCase *application.WindowUseCase) *WindowController {
    return &WindowController{
        useCase: useCase,
    }
}

func (c *WindowController) SendWindowCommand(ctx *gin.Context) {
    var command entities.WindowCommand
    if err := ctx.ShouldBindJSON(&command); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
        return
    }

    if command.State != 0 && command.State != 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "State must be 0 or 1"})
        return
    }

    err := c.useCase.SendWindowCommand(command)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send window command", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Window command sent successfully"})
}