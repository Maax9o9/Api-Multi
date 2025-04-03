package routes

import (
    "Multi/src/house/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func HouseRoutes(
    router *gin.Engine,
    addHouseController *controllers.AddHouseController,
    showAllHousesController *controllers.ShowAllHousesController,
    showHouseByIDController *controllers.ShowHouseByIDController,
    editHouseController *controllers.EditHouseController,
) {
    router.POST("/api/houses", addHouseController.AddHouse)
    router.GET("/api/houses", showAllHousesController.GetAllHouses)
    router.GET("/api/houses/:id", showHouseByIDController.GetHouseByID)
    router.PUT("/api/houses/:id", editHouseController.UpdateHouse)
}