package request

type OrganJoin struct {
	ID int `json:"id" validate:"gte=100000000,lte=999999999"`
}

type OrganAvatar struct {
	ID int `json:"id" validate:"gte=100000000,lte=999999999"`
}

type OrganSetAvatar struct {
	ID   int    `json:"id" validate:"gte=100000000,lte=999999999"`
	File string `json:"file" validate:"required"`
	Ex   string `json:"ex" validate:"oneof=jpg jpeg png"`
}

type OrganName struct {
	ID int `json:"id" validate:"gte=100000000,lte=999999999"`
}

type OrganExit struct {
	ID int `json:"id" validate:"gte=100000000,lte=999999999"`
}

type OrganList struct {
}

type OrganAvatarName struct {
	ID int `json:"id" validate:"gte=100000000,lte=999999999"`
}
