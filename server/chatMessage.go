package main

// chat msg
type ChatMessage struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
	Reciver string `json:"reciver"`
	Time string `json:"time"`
}

func (message *ChatMessage) GetContent() string {
	return message.Content
}
func (message *ChatMessage) GetReciver() string {
	return message.Reciver
}
func (message *ChatMessage) GetSender() string {
	return message.Sender
}
func (message *ChatMessage) GetTime() string {
	return message.Time
}

/* --------------------------- register  new empty message -------------------------- */
func emptyChatMessage() *ChatMessage {
	return &ChatMessage{
		Content: "",
		Sender:  "",
		Reciver: "",
		Time: "",
	}
}

