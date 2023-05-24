package request

type MessageSend struct {
	ID   int    `json:"id" validate:"len=9"`
	Data string `json:"data" validate:"required"`
}

type MessageUp struct {
	ID    int `json:"id" validate:"len=9"`
	MsgID int `json:"msgID"`
	Num   int `json:"num" validate:"required"`
}

type MessageDown struct {
	ID    int `json:"id" validate:"len=9"`
	MsgID int `json:"msgID"`
	Num   int `json:"num" validate:"required"`
}
