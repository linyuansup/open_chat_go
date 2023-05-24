package response

type MessageSend struct {
	ID int `json:"id"`
}

type MessageUp struct {
	Msg []Message `json:"message"`
}

type MessageDown struct {
	Msg []Message `json:"message"`
}
