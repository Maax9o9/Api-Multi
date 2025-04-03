package routes

import (
    "Multi/src/interruptors/gas/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func GasRoutes(
    router *gin.Engine,
    alertController *controllers.AlertGasController,
    receiveController *controllers.ReceiveGasController,
) {
    router.GET("/api/gas", receiveController.GetAllGasData)
    router.GET("/api/gas/:id", receiveController.GetGasDataByID)
    router.POST("/api/gas", alertController.CreateGasData)
}