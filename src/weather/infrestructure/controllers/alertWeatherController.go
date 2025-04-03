package controllers

import (
    "Multi/src/weather/application/services"
    "Multi/src/weather/domain/entities"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AlertWeatherController struct {
    weatherService *service.AlertWeatherService
}

func NewAlertWeatherController(weatherService *service.AlertWeatherService) *AlertWeatherController {
    return &AlertWeatherController{
        weatherService: weatherService,
    }
}

func (awc *AlertWeatherController) ProcessWeatherData(ctx *gin.Context) {
    var weatherData entities.SensorDataWeather

    if err := ctx.ShouldBindJSON(&weatherData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request format",
            "details": err.Error(),
        })
        return
    }

    err := awc.weatherService.CreateWeatherData(&weatherData)
    if err != nil {
        log.Printf("Error al guardar los datos meteorológicos: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to save weather data",
            "details": err.Error(),
        })
        return
    }

    log.Printf("Datos meteorológicos procesados y guardados: %+v", weatherData)
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Weather data processed and saved successfully",
    })
}