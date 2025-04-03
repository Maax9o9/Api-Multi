package routes

import (
    "Multi/src/interruptors/gas/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func GasRoutes(
    router *gin.Engine,
    alertController *controllers.AlertGasController,
    receiveAllController *controllers.ReceiveGasController,
    receiveByIDController *controllers.ReceiveGasByIDController,
) {
    router.GET("/api/gas", receiveAllController.GetAllGasData)
    router.GET("/api/gas/:id", receiveByIDController.GetGasDataByID)
    router.POST("/api/gas", alertController.CreateGasData)
}