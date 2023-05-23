package request

type OrganJoinRequest struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganAvatarRequest struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganSetAvatarRequest struct {
	ID   int    `json:"id" validate:"len=9"`
	File string `json:"file" validate:"required"`
	Ex   string `json:"ex" validate:"required"`
}
