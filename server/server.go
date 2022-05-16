package main

import (
	"real-time-forum/models"
	"strconv"
	"strings"
)

type WsServer struct {
	clients    map[*Client]bool // list of active
	register   chan *Client     // channel for register request
	unregister chan *Client     // channel for unreg request

	broadcast chan []byte //channel for broadcasting to all clients

	userHandler    models.UserHandlers
	messageHandler models.MessageHandlers
	postHandler    models.PostHandlers
	commentHandler models.CommentHandlers

	users []models.User //all registered users
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer(userHandler models.UserHandlers, messageHandler models.MessageHandlers, postHandler models.PostHandlers, commentHandler models.CommentHandlers) *WsServer {
	wsServer := &WsServer{
		clients:        make(map[*Client]bool),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		broadcast:      make(chan []byte),
		userHandler:    userHandler,
		messageHandler: messageHandler,
		postHandler:    postHandler,
		commentHandler: commentHandler,
	}
	// Add users from database to server
	wsServer.users = userHandler.GetAllUsers()

	return wsServer
}

/* ------------------------ function running infinit ------------------------ */
// Run our websocket server, accepting various requests
func (server *WsServer) Run() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)

		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

/* ---------- handle new client joined and notify all active users ---------- */
func (server *WsServer) registerClient(client *Client) {
	client.ID = server.userHandler.FindUserID(client)
	server.clients[client] = true //update client list
	msg := "login"
	if !contains(server.users, client) {
		msg = "signup"
		server.users = append(server.users, client) // append to all user list
	}
	server.notifyClientJoined(client, msg) // notify everybody
}

/* ----------------- helper function to check for dublicate ----------------- */
func contains(allUsers []models.User, newUser models.User) bool {
	for _, check := range allUsers {
		if check.GetName() == newUser.GetName() {
			return true
		}
	}
	return false
}

/* ----------------- log out action and notify active users ----------------- */
func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
		server.notifyClientLeft(client)
	}
}

/* ------------------- send message to all active clients ------------------- */
func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *WsServer) notifyClientJoined(client *Client, msg string) {
	message := &Message{
		Action:  UserJoinedAction,
		Sender:  client,
		Message: msg,
	}
	server.broadcastToClients(message.encode())
}

func (server *WsServer) notifyClientLeft(client *Client) {
	message := &Message{
		Action: UserLogoutAction,
		Sender: client,
	}
	server.broadcastToClients(message.encode())
}

/* --------------------------------- getinfo -------------------------------- */
func (server *WsServer) getClientStats(username string) map[string]string {
	var stats = make(map[string]string)
	posts := strconv.Itoa(server.postHandler.GetUserPostCount(username))
	comments := strconv.Itoa(server.commentHandler.GetUserCommentCount(username))
	stats["posts"] = posts       // fetch from db
	stats["comments"] = comments // fetch from db
	joinDate := server.userHandler.GetUserJoinDate(username)
	stats["joined"] = strings.Split(joinDate, "T")[0]
	return stats
}

func (server *WsServer) getChatMessages(chatId string, offset int) []models.Message {
	return server.messageHandler.GetChatMessages(chatId, offset)
}

/* ------------------ send last chat message to both users ------------------ */
func (server *WsServer) sendLastMessage(chatId string, sender models.User) {
	chatMessage := server.messageHandler.GetLastMessage(chatId) //get msg from db
	response := make([]models.Message, 1)
	response[0] = chatMessage
	// Send msg to both clients
	message := Message{
		ChatMessages: response,
		Action:       NewChatMessageAction,
		Target:       chatId,
		Sender:       sender,
		Message:      chatMessage.GetReciver(),
	}
	// send message to both parties
	for client := range server.clients {
		if client.GetName() == chatMessage.GetReciver() || client.GetName() == chatMessage.GetSender() {
			client.send <- message.encode()
		}
	}
}

/* -------------------------------------------------------------------------- */
/*                            database interactions                           */
/* -------------------------------------------------------------------------- */
func (server *WsServer) addNewMessage(msg models.Message) {
	server.messageHandler.AddMessage(msg)
}

func (server *WsServer) addNewPost(post models.Post) {
	server.postHandler.AddPost(post)
}

func (server *WsServer) findCategory(category string) string {
	return server.postHandler.SelectCategory(category)
}

func (server *WsServer) addCategory(category string) {
	server.postHandler.AddCategory(category)
}

func (server *WsServer) getAllCategories() []string {
	return server.postHandler.SelectAllCategories()
}
func (server *WsServer) getPosts(offset int) []models.Post {
	return server.postHandler.GetPosts(offset)
}
func (server *WsServer) getCategoryPosts(category string, offset int) []models.Post {
	return server.postHandler.GetCategoryPosts(category, offset)
}
func (server *WsServer) getUserPosts(offset int, username string) []models.Post {
	return server.postHandler.GetUserPosts(offset, username)
}
func (server *WsServer) getPost(postId int) models.Post {
	return server.postHandler.GetPost(postId)
}
func (server *WsServer) addNewComment(postId int, content string, author string) {
	server.commentHandler.AddComment(postId, content, author)
}

func (server *WsServer) getLastComment(postId int) models.Comment {
	return server.commentHandler.GetLastComment(postId)
}
func (server *WsServer) getComments(offset, postId int) []models.Comment {
	return server.commentHandler.GetComments(postId, offset)
}
func (server *WsServer) getUserByUsername(username string) models.User {
	return server.userHandler.FindUserByUsername(username)
}
