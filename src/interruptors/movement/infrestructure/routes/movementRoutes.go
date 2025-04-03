package routes

import (
    "Multi/src/interruptors/movement/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func MovementRoutes(
    router *gin.Engine,
    alertController *controllers.AlertMovementController,
    receiveAllController *controllers.ReceiveMovementController,
    receiveByIDController *controllers.ReceiveMovementByIDController,
) {
    router.POST("/api/movement", alertController.CreateMovementData)
    router.GET("/api/movement", receiveAllController.GetAllMovementData)
    router.GET("/api/movement/:id", receiveByIDController.GetMovementDataByID)
}