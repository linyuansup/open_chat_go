package response

type OrganJoinResponse struct{}

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
