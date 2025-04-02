package routes

import (
    "Multi/src/interruptors/light/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func LightRoutes(router *gin.Engine, controller *controllers.LightController) {
    router.GET("/api/light", controller.GetAllLightData)
    router.GET("/api/light/:id", controller.GetLightDataByID)
    router.POST("/api/light", controller.CreateLightData)
}