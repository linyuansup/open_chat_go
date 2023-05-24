package response

type FriendAgree struct{}

type FriendDisgree struct{}

type FriendRequest struct {
	ID []int `json:"id"`
}
