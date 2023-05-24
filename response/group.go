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

type GroupSetName struct{}

type GroupMember struct {
	Owner  int   `json:"owner"`
	Admin  []int `json:"admin"`
	Member []int `json:"member"`
}
