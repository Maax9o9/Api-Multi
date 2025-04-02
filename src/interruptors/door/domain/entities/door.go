package entities

import "time"

type DoorData struct {
    ID        int       `json:"id"`
    HouseID   int       `json:"house_id"`
    CreatedAt time.Time `json:"created_at"`
    Status    int       `json:"status"` 
}