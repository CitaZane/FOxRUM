package handlers

import (
	"database/sql"
	"real-time-forum/models"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
	Reciver string `json:"reciver"`
	Time    string `json:"time"`
}

func (message *Message) GetContent() string {
	return message.Content
}
func (message *Message) GetSender() string {
	return message.Sender
}
func (message *Message) GetReciver() string {
	return message.Reciver
}
func (message *Message) GetTime() string {
	return message.Time
}

/* --------------------- message- handler / db interaction --------------------- */
type MessageHandler struct {
	Db *sql.DB
}

/* --------------------------------- insert --------------------------------- */
func (handler *MessageHandler) AddMessage(message models.Message) {
	connection := ""
	if strings.Compare(message.GetSender(), message.GetReciver()) < 0 {
		connection += message.GetSender() + ":" + message.GetReciver()
	} else {
		connection += message.GetReciver() + ":" + message.GetSender()
	}
	stmt, err := handler.Db.Prepare("INSERT INTO messages(content, sender,connection) values(?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(message.GetContent(), message.GetSender(), connection)
	checkErr(err)
}

func (handler *MessageHandler) GetChatMessages(conn string, offset int) []models.Message {
	rows, err := handler.Db.Query("SELECT sender, content, created_at FROM messages WHERE connection = ? ORDER BY created_at DESC LIMIT 15 OFFSET ?", conn, offset)
	checkErr(err)
	var messages []models.Message
	defer rows.Close()
	for rows.Next() {
		var message Message
		var dateTime time.Time
		rows.Scan(&message.Sender, &message.Content, &dateTime)
		message.Time = strconv.FormatInt(dateTime.Unix(), 10)
		messages = append(messages, &message)
	}
	return messages
}
func (handler *MessageHandler) GetLastMessage(conn string) models.Message {
	row := handler.Db.QueryRow("SELECT sender, content, created_at FROM messages WHERE connection = ? ORDER BY created_at DESC LIMIT 1", conn)
	var message models.Message
	var temp Message
	var dateTime time.Time
	if err := row.Scan(&temp.Sender, &temp.Content, &dateTime); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}
	splitConn := strings.Split(conn, ":")
	if splitConn[0] == temp.Sender {
		temp.Reciver = splitConn[1]
	} else {
		temp.Reciver = splitConn[0]
	}
	temp.Time = strconv.FormatInt(dateTime.Unix(), 10)
	message = &temp
	return message
}
