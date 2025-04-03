package routes

import (
    "Multi/src/weather/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WeatherRoutes(
    router *gin.Engine,
    alertController *controllers.AlertWeatherController,
    receiveController *controllers.ReceiveWeatherController,
) {
    router.GET("/api/weather", receiveController.GetAllWeatherData)
    router.GET("/api/weather/latest", receiveController.GetLatestWeatherData)
    router.POST("/api/weather", alertController.ProcessWeatherData)
}