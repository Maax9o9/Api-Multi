package entities

type HouseProfile struct {
	HouseID      int    `json:"house_id"`
	OwnerID      int    `json:"owner_id"` // Renombrado de UserID a OwnerID para mayor claridad
	UbicationGps string `json:"ubication_gps"`
	HouseName    string `json:"house_name"`
	Image        string `json:"image"`
	DeviceCode   string `json:"device_code"` // Código de la ESP32
	CreatedAt    int64  `json:"created_at"`
}
