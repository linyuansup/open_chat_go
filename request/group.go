package request

type GroupCreate struct {
	Name string `json:"name" validate:"required"`
}

type GroupDelete struct {
	ID int `json:"id" validate:"gte=100000000,lte=599999999"`
}

type GroupAgree struct {
	GroupID int `json:"groupID" validate:"gte=600000000,lte=999999999"`
	UserID  int `json:"userID" validate:"gte=100000000,lte=599999999"`
}

type GroupSetAdmin struct {
	GroupID int `json:"groupID" validate:"gte=600000000,lte=999999999"`
	UserID  int `json:"userID" validate:"gte=100000000,lte=599999999"`
}

type GroupRemoveAdmin struct {
	GroupID int `json:"groupID" validate:"gte=600000000,lte=999999999"`
	UserID  int `json:"userID" validate:"gte=100000000,lte=599999999"`
}

type GroupRequest struct{}

type GroupDisagree struct {
	GroupID int `json:"groupID" validate:"gte=600000000,lte=999999999"`
	UserID  int `json:"userID" validate:"gte=100000000,lte=599999999"`
}

type GroupSetName struct {
	ID   int    `json:"groupID" validate:"gte=600000000,lte=999999999"`
	Name string `json:"name" validate:"required"`
}

type GroupMember struct {
	ID int `json:"groupID" validate:"gte=600000000,lte=999999999"`
}
