package routes

import (
    "Multi/src/house/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func HouseRoutes(router *gin.Engine, addHouseController *controllers.AddHouseController, showHouseController *controllers.ShowHouseController, editHouseController *controllers.EditHouseController) {
    router.POST("/api/houses", addHouseController.AddHouse)
    router.GET("/api/houses", showHouseController.GetAllHouses)
    router.GET("/api/houses/:id", showHouseController.GetHouseByID)
    router.PUT("/api/houses/:id", editHouseController.UpdateHouse)
}