package response

type OrganJoin struct{}

type OrganAvatar struct {
	File string `json:"file"`
	Ex   string `json:"ex"`
}

type OrganSetAvatar struct {
	Name string `json:"name"`
}

type OrganName struct {
	Name string `json:"name"`
}

type OrganExit struct {
}

type OrganList struct {
	ID []int `json:"id"`
}

type OrganAvatarName struct {
	Name string `json:"name"`
}
