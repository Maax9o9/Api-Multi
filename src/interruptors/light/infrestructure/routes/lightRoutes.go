package routes

import (
    "Multi/src/interruptors/light/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func LightRoutes(
    router *gin.Engine,
    alertController *controllers.AlertLightController,
    receiveController *controllers.ReceiveLightController,
) {
    router.GET("/api/light", receiveController.GetAllLightData)
    router.GET("/api/light/:id", receiveController.GetLightDataByID)
    router.POST("/api/light", alertController.CreateLightData)
}