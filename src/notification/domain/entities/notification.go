package entities

import "time"

type Notification struct {
    ID         int       `json:"id"`
    SensorID   int       `json:"sensor_id"`
    SensorType string    `json:"sensor_type"`
    Date       time.Time `json:"date"`
    Message    string    `json:"message"`
}