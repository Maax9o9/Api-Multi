package routes

import (
    "Multi/src/weather/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WeatherRoutes(
    router *gin.Engine,
    alertController *controllers.AlertWeatherController,
    receiveAllController *controllers.ReceiveWeatherController,
    receiveByIDController *controllers.ReceiveWeatherByIDController,
) {
    router.GET("/api/weather", receiveAllController.GetAllWeatherData)
    router.GET("/api/weather/latest", receiveAllController.GetLatestWeatherData)
    router.GET("/api/weather/:id", receiveByIDController.GetWeatherDataByID)
    router.POST("/api/weather", alertController.ProcessWeatherData)
}