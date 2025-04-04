package controllers

import (
    "bytes"
    "encoding/json"
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

    go func() {
        url := "http://localhost:7070/weather"
        payload := map[string]interface{}{
            "weather_id": weatherData.WeatherID,
            "date":       weatherData.Date,
            "heat":       weatherData.Heat,
            "damp":       weatherData.Damp,
        }
        jsonData, _ := json.Marshal(payload)

        println("Sending weather data to WebSocket:")
        println("URL:", url)
        println("Payload:", string(jsonData))

        resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            println("Failed to send weather data:", err.Error())
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            println("Weather data sent successfully!")
        } else {
            println("Failed to send weather data. Status code:", resp.StatusCode)
        }
    }()

    log.Printf("Datos meteorológicos procesados y guardados: %+v", weatherData)
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Weather data processed and saved successfully",
    })
}