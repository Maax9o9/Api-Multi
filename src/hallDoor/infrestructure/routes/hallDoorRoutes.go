package routes

import (
    "Multi/src/hallDoor/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func HallDoorRoutes(
    router *gin.Engine,
    alertController *controllers.AlertHallDoorController,
    receiveAllController *controllers.ReceiveHallDoorController,
    receiveByIDController *controllers.ReceiveHallDoorByIDController,
    updateController *controllers.UpdateHallDoorController,
) {
    router.GET("/api/hallDoor", receiveAllController.GetAllHallDoors)
    router.GET("/api/hallDoor/latest", receiveAllController.GetLatestHallDoorData)
    router.GET("/api/hallDoor/:id", receiveByIDController.GetHallDoorByID)
    router.POST("/api/hallDoor", alertController.ProcessHallDoorData)
    router.PUT("/api/hallDoor/:id", updateController.UpdateHallDoorStatus)
}