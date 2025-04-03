package routes

import (
    "Multi/src/interruptors/door/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func DoorRoutes(
    router *gin.Engine,
    alertController *controllers.AlertDoorController,
    getAllController *controllers.GetAllDoorController,
    getByIDController *controllers.GetDoorByIDController,
) {
    router.POST("/api/door", alertController.CreateDoorData)
    router.GET("/api/door", getAllController.GetAllDoorData)
    router.GET("/api/door/:id", getByIDController.GetDoorDataByID)
}