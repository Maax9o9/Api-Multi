package entities

import "time"

type HallDoor struct {
    ID     int       `json:"id"`
    Date   time.Time `json:"date"`
    Status int       `json:"status"` 
}