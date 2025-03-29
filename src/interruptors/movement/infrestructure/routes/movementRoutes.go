package routes

import (
    "Multi/src/interruptors/movement/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func MovementRoutes(router *gin.Engine, controller *controllers.MovementController) {
    router.POST("/api/movement", controller.SendMovementCommand)
}