package types

type Message struct {
	FromUser int    `json:"fromUser,string"`
	ToUser   int    `json:"toUser,string"`
	Message  string `json:"message"`
}

type MessageRequest struct {
	FromUser int `json:"fromUser,string"`
	ToUser   int `json:"toUser,string"`
}
