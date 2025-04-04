package entities

type HouseMember struct {
	ID       int    `json:"id"`
	HouseID  int    `json:"house_id"`
	UserID   int    `json:"user_id"`
	Role     string `json:"role"` // "owner", "admin", "member"
	JoinedAt int64  `json:"joined_at"`
}
