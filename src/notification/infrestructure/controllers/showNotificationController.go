package controllers

import (
    "Multi/src/notification/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ShowAllNotificationsController struct {
    useCase *application.ShowAllNotificationsUseCase
}

func NewShowAllNotificationsController(useCase *application.ShowAllNotificationsUseCase) *ShowAllNotificationsController {
    return &ShowAllNotificationsController{
        useCase: useCase,
    }
}

func (c *ShowAllNotificationsController) GetAllNotifications(ctx *gin.Context) {
    notifications, err := c.useCase.GetAllNotifications()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, notifications)
}