package model

type Message struct {
	Sender     string    `json:"sender"`
	Recipient  Recipient `json:"recipient"`
	Text       string    `json:"text"`
	Attachment string    `json:"attachment"`
}

type Recipient struct {
	IsGroup   bool     `json:"isGroup"`
	GroupID   string   `json:"groupID"`
	Receivers []string `json:"receivers"`
}
