package controllers

import (
    "Multi/src/notification/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ShowNotificationController struct {
    useCase *application.ShowNotificationUseCase
}

func NewShowNotificationController(useCase *application.ShowNotificationUseCase) *ShowNotificationController {
    return &ShowNotificationController{
        useCase: useCase,
    }
}

func (c *ShowNotificationController) GetAllNotifications(ctx *gin.Context) {
    notifications, err := c.useCase.GetAllNotifications()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, notifications)
}

func (c *ShowNotificationController) GetNotificationByID(ctx *gin.Context) {
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