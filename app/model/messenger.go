package model

type Message struct {
	Message     string   `json:"message"`
	Sender      string   `json:"number"`
	Recipients  []string `json:"recipients"`
	Attachments []string `json:"base64_attachments"`
}

type MessengerClient interface {
	Push(holding Holding, message Message) error
}
