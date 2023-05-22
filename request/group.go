package request

type GroupCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type GroupDeleteRequest struct {
	ID int `json:"id" validate:"len=9"`
}

type GroupAgreeRequest struct {
	GroupID int `json:"groupID" validate:"len=9"`
	UserID  int `json:"userID" validate:"len=9"`
}

type GroupSetAdminRequest struct {
	GroupID int `json:"groupID" validate:"len=9"`
	UserID  int `json:"userID" validate:"len=9"`
}
