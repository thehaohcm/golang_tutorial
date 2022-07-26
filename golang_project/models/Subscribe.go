package models

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type SubscribeResponse struct {
	Success bool `json:"success"`
}

type BlockSubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type BlockSubscribeResponse struct {
	Success bool `json:"success"`
}

type GetSubscribingEmailListRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type GetSubscribingEmailListResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
