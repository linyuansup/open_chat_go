package response

type OrganJoinResponse struct{}

type OrganAvatarResponse struct {
	File string `json:"file"`
	Ex   string `json:"ex"`
}

type OrganSetAvatarResponse struct{}
