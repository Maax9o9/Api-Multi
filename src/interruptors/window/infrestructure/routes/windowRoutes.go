package routes

import (
    "Multi/src/interruptors/window/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WindowRoutes(
    router *gin.Engine,
    alertController *controllers.AlertWindowController,
    receiveAllController *controllers.ReceiveWindowController,
    receiveByIDController *controllers.ReceiveWindowByIDController,
) {
    router.GET("/api/window", receiveAllController.GetAllWindowData)
    router.GET("/api/window/:id", receiveByIDController.GetWindowDataByID)
    router.POST("/api/window", alertController.CreateWindowData)
}