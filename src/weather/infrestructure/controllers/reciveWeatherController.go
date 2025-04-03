package controllers

import (
    "Multi/src/weather/application/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ReceiveWeatherController struct {
    weatherService *service.ReceiveWeatherService
}

func NewReceiveWeatherController(weatherService *service.ReceiveWeatherService) *ReceiveWeatherController {
    return &ReceiveWeatherController{
        weatherService: weatherService,
    }
}

func (rwc *ReceiveWeatherController) GetAllWeatherData(c *gin.Context) {
    data, err := rwc.weatherService.GetAllWeatherData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Error al obtener los datos meteorológicos",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}

func (rwc *ReceiveWeatherController) GetLatestWeatherData(c *gin.Context) {
    data, err := rwc.weatherService.GetLatestWeatherData()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "No hay datos meteorológicos recientes",
            "details": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, data)
}