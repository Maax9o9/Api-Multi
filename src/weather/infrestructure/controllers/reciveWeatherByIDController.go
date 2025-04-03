package controllers

import (
    "Multi/src/weather/application/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReceiveWeatherByIDController struct {
    weatherService *service.ReceiveWeatherService
}

func NewReceiveWeatherByIDController(weatherService *service.ReceiveWeatherService) *ReceiveWeatherByIDController {
    return &ReceiveWeatherByIDController{
        weatherService: weatherService,
    }
}

func (rwc *ReceiveWeatherByIDController) GetWeatherDataByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "ID inválido",
        })
        return
    }

    data, err := rwc.weatherService.GetWeatherDataByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "Error al obtener los datos meteorológicos",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, data)
}