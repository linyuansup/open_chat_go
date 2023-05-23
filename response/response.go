package response

type Response[T any] struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

type Request struct {
	ID      int `json:"id"`
	GroupID int `json:"groupID"`
}
