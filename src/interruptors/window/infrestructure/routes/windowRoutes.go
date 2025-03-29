package routes

import (
    "Multi/src/interruptors/window/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WindowRoutes(router *gin.Engine, controller *controllers.WindowController) {
    router.POST("/api/window", controller.SendWindowCommand)
}