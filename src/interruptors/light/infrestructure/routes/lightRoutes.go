package routes

import (
    "Multi/src/interruptors/light/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func LightRoutes(router *gin.Engine, controller *controllers.LightController) {
    router.POST("/api/light", controller.SendLightCommand)
}