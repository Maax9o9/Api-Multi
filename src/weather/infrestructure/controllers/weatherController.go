package controllers

import (
    "net/http"
    "Multi/src/weather/application/services"
    "github.com/gin-gonic/gin"
)

type WeatherController struct {
    weatherService *service.WeatherService
}

func NewWeatherController(weatherService *service.WeatherService) *WeatherController {
    return &WeatherController{
        weatherService: weatherService,
    }
}

func (wc *WeatherController) GetAllWeatherData(c *gin.Context) {
    data, err := wc.weatherService.GetAllWeatherData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos meteorológicos", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (wc *WeatherController) GetLatestWeatherData(c *gin.Context) {
    data, err := wc.weatherService.GetLatestWeatherData()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No hay datos meteorológicos recientes", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}