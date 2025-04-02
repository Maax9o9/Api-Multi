package routes

import (
    "Multi/src/interruptors/window/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WindowRoutes(router *gin.Engine, controller *controllers.WindowController) {
    router.GET("/api/window", controller.GetAllWindowData)
    router.GET("/api/window/:id", controller.GetWindowDataByID)
    router.POST("/api/window", controller.CreateWindowData)
}