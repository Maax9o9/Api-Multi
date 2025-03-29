package entities

import "time"

type Notification struct {
	ID               int       `json:"id"`
	HouseID          int       `json:"house_id"`
	Date             time.Time `json:"date"`
	Message          string    `json:"message"`
	TypeNotification string    `json:"type_notification"`
}