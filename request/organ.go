package request

type OrganJoinRequest struct {
	ID int `json:"id" validate:"len=9"`
}
