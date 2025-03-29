package routes

import (
    "Multi/src/notification/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func NotificationRoutes(router *gin.Engine, createController *controllers.CreateNotificationController, showController *controllers.ShowNotificationController) {
    router.POST("/api/notifications", createController.CreateNotification)
    router.GET("/api/notifications", showController.GetAllNotifications)
    router.GET("/api/notifications/:id", showController.GetNotificationByID)
}