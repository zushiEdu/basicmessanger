package types

type Message struct {
	ToUser  int    `json:"toUser,string"`
	Message string `json:"message"`
}

type MessageRequest struct {
	InvolvingUser int `json:"involving"`
	Token         string
}

type MessageResponse struct {
	Message  string `json:"message"`
	UserFrom int    `json:"userFrom"`
}
