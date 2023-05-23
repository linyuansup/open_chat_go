package request

type OrganJoinRequest struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganAvatarRequest struct {
	ID int `json:"id" validate:"len=9"`
}
