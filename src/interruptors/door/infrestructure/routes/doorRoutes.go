package routes

import (
    "Multi/src/interruptors/door/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func DoorRoutes(
    router *gin.Engine,
    alertController *controllers.AlertDoorController,
    receiveController *controllers.ReceiveDoorController,
) {
    router.POST("/api/door", alertController.CreateDoorData)
    router.GET("/api/door", receiveController.GetAllDoorData)
    router.GET("/api/door/:id", receiveController.GetDoorDataByID)
}