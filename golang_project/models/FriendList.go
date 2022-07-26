package models

type FriendListRequest struct{
	Email string `json:"email"`
}

type FriendListResponse struct{
	Success bool `json:"success"`
	Friends []string `json:"friends"`
	Count int `json:"count"`
}