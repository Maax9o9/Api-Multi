package routes

import (
    "Multi/src/interruptors/window/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WindowRoutes(
    router *gin.Engine,
    alertController *controllers.AlertWindowController,
    receiveController *controllers.ReceiveWindowController,
) {
    router.GET("/api/window", receiveController.GetAllWindowData)
    router.GET("/api/window/:id", receiveController.GetWindowDataByID)

    router.POST("/api/window", alertController.CreateWindowData)
}