package response

type UserCreate struct {
	ID uint `json:"id"`
}

type UserLogin struct {
	ID uint `json:"id"`
}

type UserSetPassword struct {
}

type UserSetName struct {
}
