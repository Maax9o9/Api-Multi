package entities

import "time"

type GasSensor struct {
    ID        int       `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Status    int       `json:"status"` 
    GasLevel  float64   `json:"gas_level"` 
}