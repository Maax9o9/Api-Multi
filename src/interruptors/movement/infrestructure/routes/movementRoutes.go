package routes

import (
    "Multi/src/interruptors/movement/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func MovementRoutes(
    router *gin.Engine,
    alertController *controllers.AlertMovementController,
    receiveController *controllers.ReceiveMovementController,
) {
    router.POST("/api/movement", alertController.CreateMovementData)
    router.GET("/api/movement", receiveController.GetAllMovementData)
    router.GET("/api/movement/:id", receiveController.GetMovementDataByID)
}