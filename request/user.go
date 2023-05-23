package request

type UserCreate struct {
	PhoneNumber string `json:"phoneNumber" validate:"number,len=11"`
	Password    string `json:"password" validate:"len=32"`
	DeviceID    string `json:"deviceID" validate:"required"`
}

type UserLogin struct {
	PhoneNumber string `json:"phoneNumber" validate:"number,len=11"`
	Password    string `json:"password" validate:"len=32"`
	DeviceID    string `json:"deviceID" validate:"required"`
}

type UserSetPassword struct {
	OldPassword string `json:"oldPassword" validate:"len=32"`
	Password    string `json:"password" validate:"len=32"`
}

type UserSetName struct {
	Name string `json:"name" validate:"required"`
}
