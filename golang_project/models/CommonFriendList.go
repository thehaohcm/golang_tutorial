package models

type CommonFriendListRequest struct{
	Friends []string `json:"friends"`
}

type CommonFriendListResponse struct{
	Success bool `json:"success"`
	Friends []string `json:"friends"`
	Count int `json:"count"`
}