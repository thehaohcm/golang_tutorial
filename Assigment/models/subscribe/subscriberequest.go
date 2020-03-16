package models

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
