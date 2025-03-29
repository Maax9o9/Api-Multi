package entities

type HouseProfile struct {
	HouseID      int    `json:"house_id"`
	UserID       int    `json:"user_id"`
	UbicationGps string `json:"ubication_gps"`
	HouseName    string `json:"house_name"`
	Image        string `json:"image"`
}