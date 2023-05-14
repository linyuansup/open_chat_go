package response

type Response[T any] struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}
