package routes

import (
    "Multi/src/notification/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func NotificationRoutes(
    router *gin.Engine,
    createController *controllers.CreateNotificationController,
    showAllController *controllers.ShowAllNotificationsController,
    showByIDController *controllers.ShowNotificationByIDController,
) {
    router.POST("/api/notifications", createController.CreateNotification)
    router.GET("/api/notifications", showAllController.GetAllNotifications)
    router.GET("/api/notifications/:id", showByIDController.GetNotificationByID)
}