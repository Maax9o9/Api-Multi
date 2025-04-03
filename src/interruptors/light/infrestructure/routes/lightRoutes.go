package routes

import (
    "Multi/src/interruptors/light/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func LightRoutes(
    router *gin.Engine,
    alertController *controllers.AlertLightController,
    receiveAllController *controllers.ReceiveLightController,
    receiveByIDController *controllers.ReceiveLightByIDController,
) {
    router.GET("/api/light", receiveAllController.GetAllLightData)
    router.GET("/api/light/:id", receiveByIDController.GetLightDataByID)
    router.POST("/api/light", alertController.CreateLightData)
}