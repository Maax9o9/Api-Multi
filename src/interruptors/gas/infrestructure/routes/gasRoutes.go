package routes

import (
    "Multi/src/interruptors/gas/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func GasRoutes(router *gin.Engine, controller *controllers.GasController) {
    router.GET("/api/gas", controller.GetAllGasData)
    router.GET("/api/gas/:id", controller.GetGasDataByID)
    router.POST("/api/gas", controller.CreateGasData)
}