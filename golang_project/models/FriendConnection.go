package models

type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

type FriendConnectionResponse struct {
	Success bool `json:"success"`
}
