package request

type GroupCreateRequest struct {
	Name string `json:"name" validate:"required"`
}