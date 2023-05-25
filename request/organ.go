package request

type OrganJoin struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganAvatar struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganSetAvatar struct {
	ID   int    `json:"id" validate:"len=9"`
	File string `json:"file" validate:"required"`
	Ex   string `json:"ex" validate:"required"`
}

type OrganName struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganExit struct {
	ID int `json:"id" validate:"len=9"`
}

type OrganList struct {
}

type OrganAvatarName struct {
	ID int `json:"id" validate:"len=9"`
}
