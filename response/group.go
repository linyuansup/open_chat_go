package response

type GroupCreate struct {
	ID int `json:"id"`
}

type GroupDelete struct{}

type GroupAgree struct{}

type GroupSetAdmin struct{}

type GroupRemoveAdmin struct{}
