package routes

import (
    "Multi/src/weather/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func WeatherRoutes(router *gin.Engine, weatherController *controllers.WeatherController) {
    router.GET("/api/weather", weatherController.GetAllWeatherData)
    router.GET("/api/weather/latest", weatherController.GetLatestWeatherData)
}