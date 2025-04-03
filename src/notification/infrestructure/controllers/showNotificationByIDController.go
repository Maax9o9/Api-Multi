package controllers

import (
    "Multi/src/notification/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ShowNotificationByIDController struct {
    useCase *application.ShowNotificationByIDUseCase
}

func NewShowNotificationByIDController(useCase *application.ShowNotificationByIDUseCase) *ShowNotificationByIDController {
    return &ShowNotificationByIDController{
        useCase: useCase,
    }
}

func (c *ShowNotificationByIDController) GetNotificationByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    notification, err := c.useCase.GetNotificationByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Notification not found", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, notification)
}