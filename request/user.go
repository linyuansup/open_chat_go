package request

type UserCreateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"number,len=11"`
	Password    string `json:"password" validate:"len=32"`
	DeviceID    string `json:"deviceID" validate:"required"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"number,len=11"`
	Password    string `json:"password" validate:"len=32"`
	DeviceID    string `json:"deviceID" validate:"required"`
}

type UserSetPasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"len=32"`
	Password    string `json:"password" validate:"len=32"`
}
