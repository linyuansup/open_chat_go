package request

type FriendAgree struct {
	ID int `json:"id" validate:"gte=100000000,lte=599999999"`
}

type FriendDisgree struct {
	ID int `json:"id" validate:"gte=100000000,lte=599999999"`
}

type FriendRequest struct{}
