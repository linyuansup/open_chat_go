package response

type GroupCreate struct {
	ID int `json:"id"`
}

type GroupDelete struct{}

type GroupAgree struct{}

type GroupSetAdmin struct{}

type GroupRemoveAdmin struct{}

type GroupDisagree struct{}

type GroupRequest struct {
	Request []Request `json:"request"`
}
