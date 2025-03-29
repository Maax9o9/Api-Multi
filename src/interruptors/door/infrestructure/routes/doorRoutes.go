package routes

import (
    "Multi/src/interruptors/door/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func DoorRoutes(router *gin.Engine, controller *controllers.DoorController) {
    router.POST("/api/door", controller.SendDoorCommand)
}