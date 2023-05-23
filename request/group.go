package request

type GroupCreate struct {
	Name string `json:"name" validate:"required"`
}

type GroupDelete struct {
	ID int `json:"id" validate:"len=9"`
}

type GroupAgree struct {
	GroupID int `json:"groupID" validate:"len=9"`
	UserID  int `json:"userID" validate:"len=9"`
}

type GroupSetAdmin struct {
	GroupID int `json:"groupID" validate:"len=9"`
	UserID  int `json:"userID" validate:"len=9"`
}

type GroupRemoveAdmin struct {
	GroupID int `json:"groupID" validate:"len=9"`
	UserID  int `json:"userID" validate:"len=9"`
}

type GroupRequest struct{}
