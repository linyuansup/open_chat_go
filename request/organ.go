package request

type GetAvatarRequest struct {
	ID int `json:"id" validate:"len=9"`
}