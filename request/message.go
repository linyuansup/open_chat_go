package request

type MessageSend struct {
	ID   int    `json:"id" validate:"len=9"`
	Data string `json:"data" validate:"required"`
}
