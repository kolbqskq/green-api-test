package client

type SendMessageResponse struct {
	StatusCode int
	MessageID  string `json:"idMessage"`
	Error      *APIError
}

type GetChatHistoryRequest struct {
	ChatID string
}

type GetChatHistoryResponse struct {
	StatusCode int
	Messages   []Message
	Error      *APIError
}

type Message struct {
	Type        string `json:"type"`
	MessageID   string `json:"idMessage"`
	Timestamp   int64  `json:"timestamp"`
	TypeMessage string `json:"typeMessage"`
	ChatID      string `json:"chatId"`
	TextMessage string `json:"textMessage"`
}

type GetStateInstanceResponse struct {
	StatusCode    int
	StateInstance string `json:"stateInstance"`
	Error         *APIError
}

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	Path       string `json:"path"`
}
