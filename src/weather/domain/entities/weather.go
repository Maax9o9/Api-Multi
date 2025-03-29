package entities
import "time"

type SensorDataWeather struct {
	WeatherID int       `json:"weather_id"`
	HouseID   int       `json:"house_id"`
	Date      time.Time `json:"date"`
	Heat      float64   `json:"heat"`
	Damp      float64   `json:"damp"`
}