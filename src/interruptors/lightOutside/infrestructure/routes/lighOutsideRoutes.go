package routes

import (
    "Multi/src/interruptors/lightOutside/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func LightRoutes(
    router *gin.Engine,
    alertController *controllers.AlertLightController,
    receiveAllController *controllers.ReceiveLightController,
    receiveByIDController *controllers.ReceiveLightByIDController,
) {
    router.GET("/api/lightOutside", receiveAllController.GetAllLightData)
    router.GET("/api/lightOutside/:id", receiveByIDController.GetLightDataByID)
    router.POST("/api/lightOutside", alertController.CreateLightData)
}