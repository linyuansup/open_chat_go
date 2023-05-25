package request

type MessageSend struct {
	ID   int    `json:"id" validate:"gte=100000000,lte=999999999"`
	Data string `json:"data" validate:"required"`
}

type MessageUp struct {
	ID    int `json:"id" validate:"gte=100000000,lte=999999999"`
	MsgID int `json:"msgID" validate:"required"`
	Num   int `json:"num" validate:"required"`
}

type MessageDown struct {
	ID    int `json:"id" validate:"gte=100000000,lte=999999999"`
	MsgID int `json:"msgID" validate:"required"`
	Num   int `json:"num" validate:"required"`
}
