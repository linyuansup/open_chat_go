package request

type GroupCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type GroupDeleteRequest struct {
	ID int `json:"id" validate:"len=9"`
}
