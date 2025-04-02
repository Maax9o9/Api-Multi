package routes

import (
    "Multi/src/interruptors/door/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func DoorRoutes(router *gin.Engine, controller *controllers.DoorController) {
    router.POST("/api/door", controller.CreateDoorData)
    router.GET("/api/door", controller.GetAllDoorData)
    router.GET("/api/door/:id", controller.GetDoorDataByID)
}