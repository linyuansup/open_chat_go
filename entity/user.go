package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DeviceID string `json:"deviceID"`
}
