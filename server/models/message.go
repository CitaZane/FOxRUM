package models

/* ------------------------ register function on message----------------------- */
type Message interface {
	GetContent() string
	GetReciver() string
	GetSender() string
	GetTime() string
}
/* ------------- register functions that interact with database ------------- */
type MessageHandlers interface {
	AddMessage(message Message)
	GetChatMessages(conn string, offset int) []Message
	GetLastMessage(conn string) Message
}