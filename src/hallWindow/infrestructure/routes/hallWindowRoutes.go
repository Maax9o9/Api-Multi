package routes

import (
    "Multi/src/hallWindow/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func HallWindowRoutes(
    router *gin.Engine,
    alertController *controllers.AlertHallWindowController,
    receiveAllController *controllers.ReceiveHallWindowController,
    receiveByIDController *controllers.ReceiveHallWindowByIDController,
    updateController *controllers.UpdateHallWindowController,
) {
    router.GET("/api/hallWindow", receiveAllController.GetAllHallWindows)
    router.GET("/api/hallWindow/latest", receiveAllController.GetLatestHallWindowData)
    router.GET("/api/hallWindow/:id", receiveByIDController.GetHallWindowByID)
    router.POST("/api/hallWindow", alertController.ProcessHallWindowData)
    router.PUT("/api/hallWindow/:id", updateController.UpdateHallWindowStatus)
}