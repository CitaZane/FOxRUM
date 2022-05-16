package main

import (
	"encoding/json"
	"log"
	"net/http"
	"real-time-forum/auth"
	"real-time-forum/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second
	// Max time till next pong from peer
	pongWait = 60 * time.Second
	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	// space   = []byte{' '}
)

//  upgrader is used to upgrade the HTTP server connection to the WebSocket protocol.
var upgrader = websocket.Upgrader{}

// Client represents the websocket client at the server
type Client struct {
	Name        string `json:"name"`
	ID          string `json:"-"`
	Online      bool   `json:"online"`
	LastMsgTime string `json:"lastMsgTime"`
	ActiveChats map[string]bool
	conn        *websocket.Conn // The actual websocket connection.
	wsServer    *WsServer       //pointer to chat socket server
	send        chan []byte
}

/* --------------------------- register new client -------------------------- */
func newClient(conn *websocket.Conn, wsServer *WsServer, name string) *Client {
	return &Client{
		Name:        name,
		Online:      true,
		ActiveChats: make(map[string]bool),
		conn:        conn,
		wsServer:    wsServer,
		send:        make(chan []byte, 256),
	}
}

func (client *Client) GetName() string {
	return client.Name
}
func (client *Client) GetId() string {
	return client.ID
}

func (client *Client) GetPassword() string {
	return ""
}
func (client *Client) IsOnline() {
	client.Online = true
}
func (client *Client) IsOffline() {
	client.Online = false
}
func (client *Client) SetMessageTime(time string) {
	client.LastMsgTime = time
}

/* ----------------- handle all incoming messages for client ---------------- */
func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	/* ----------------- creating message-> configuring actions ----------------- */
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}
	/* --------- Attach the client object as the sender of the messsage. -------- */
	message.Sender = client
	/* -------------------- find correct handler for message -------------------- */
	switch message.Action {
	case GetAllUsersAction:
		client.handleGetAllUsers()
	case GetUsernameAction:
		client.handleGetUsername()
	case UserLogoutAction:
		client.handleLogout(message)
	case GetUserStatAction:
		client.handleGetUserStats()
	case NewChatMessageAction:
		client.handleNewChatMessage(message)
	case JoinChatAction:
		client.handleJoinChat(message)
	case ChatMessageRequestAction:
		client.handleChatMessageRequest(message)
	case NewPostAction:
		client.handleNewPost(message)
	case GetCategoriesAction:
		client.handleGetCategories()
	case GetPostsAction:
		client.handleGetPosts(message)
	case GetCategoryPostsAction:
		client.handleGetCategoryPosts(message)
	case GetUserPostsAction:
		client.handleGetUserPosts(message)
	case GetPostAction:
		client.handleGetPost(message)
	case NewCommentAction:
		client.handleNewComment(message)
	case GetCommentsAction:
		client.handleGetComments(message)
	case TypeAction:
		client.handleType(message)
	}
}

/* -------- construct user list from db and registered online clinets ------- */
func (client *Client) handleGetAllUsers() {
	allUsers := []models.User{}

	for _, user := range client.wsServer.users {
		if client.GetName() == user.GetName() {
			continue
		}
		user.IsOffline()
		// check if user online -> turn on online flag
		for onlineUser := range client.wsServer.clients {
			if onlineUser.GetName() == user.GetName() {
				user.IsOnline()
				break
			}
		}
		chatId := getChatId(client.GetName(), user.GetName())
		chatMessage := client.wsServer.messageHandler.GetLastMessage(chatId)
		if chatMessage != nil {
			user.SetMessageTime(chatMessage.GetTime())
		} else {
			user.SetMessageTime("")
		}
		allUsers = append(allUsers, user)
	}
	message := Message{
		Users:  allUsers,
		Action: GetAllUsersAction,
	}
	client.send <- message.encode()
}

func (client *Client) handleGetUsername() {
	message := Message{
		Message: client.GetName(),
		Action:  GetUsernameAction,
	}
	client.send <- message.encode()
}

func (client *Client) handleGetUserStats() {
	stats := client.wsServer.getClientStats(client.GetName())
	message := Message{
		Stats:  stats,
		Action: GetUserStatAction,
	}
	client.send <- message.encode()
}

/* ----------------------- responses to actions invoked ---------------------- */
func (client *Client) handleLogout(message Message) {
	for activeClient := range client.wsServer.clients {
		if activeClient.GetName() == client.GetName() {
			message := &Message{
				Action: ThrowOutAction,
				Sender: activeClient,
			}
			activeClient.send <- message.encode()
			activeClient.wsServer.unregister <- activeClient
		}
	}
}

/* ------------------------- unregister from server ------------------------- */
func (client *Client) disconnect() {
	client.wsServer.unregister <- client // unregister
}

/* -------------------------------- chatting -------------------------------- */
func (client *Client) handleNewChatMessage(incoming Message) {
	// add message into database
	// send message to both parties
	msg := emptyChatMessage()
	json.Unmarshal([]byte(incoming.Message), msg)
	client.wsServer.addNewMessage(msg)
	//Respond to client back
	chatId := getChatId(msg.GetSender(), msg.GetReciver()) // construct chatId
	client.wsServer.sendLastMessage(chatId, incoming.Sender)
}

// Client joining (clicking) on chat
// if chat not active already, then register
func (client *Client) handleJoinChat(msg Message) {
	reciver := client.wsServer.getUserByUsername(msg.Target)
	response := ""
	if reciver != nil {
		response = "valid"
	}
	chatId := getChatId(client.Name, msg.Target) // construct chatId
	if !client.ActiveChats[chatId] {             // check if chat already active?
		// if not active -> register open chat for client
		client.ActiveChats[chatId] = true
	}
	chatHistory := client.wsServer.getChatMessages(chatId, 0)
	message := Message{
		ChatMessages: chatHistory,
		Action:       JoinChatAction,
		Target:       chatId,
		Message:      response,
	}
	client.send <- message.encode()
}

func (client *Client) handleChatMessageRequest(msg Message) {
	offset, _ := strconv.Atoi(msg.Message)
	chatId := getChatId(client.Name, msg.Target)
	chatHistory := client.wsServer.getChatMessages(chatId, offset)
	message := Message{
		ChatMessages: chatHistory,
		Action:       ChatMessageRequestAction,
		Target:       chatId,
	}
	client.send <- message.encode()
}
func (client *Client) handleType(msg Message) {
	chatId := getChatId(client.Name, msg.Target)
	message := Message{
		Message: msg.Message,
		Action:  TypeAction,
		Target:  chatId,
		Sender:  client,
	}
	for user := range client.wsServer.clients {
		if user.GetName() == msg.Target {
			user.send <- message.encode()
		}
	}
}

func (client *Client) handleNewPost(incoming Message) {
	post := postTemplate()
	json.Unmarshal([]byte(incoming.Message), post)
	if post.valid() {
		var categoryDB = client.wsServer.findCategory(post.GetCategory())
		if categoryDB == "" {
			client.wsServer.addCategory(post.GetCategory())
		}
		client.wsServer.addNewPost(post)
	}
}

func (client *Client) handleGetCategories() {
	categories := client.wsServer.getAllCategories()
	message := Message{
		Action:     GetCategoriesAction,
		Categories: categories,
	}
	client.send <- message.encode()
}

func (client *Client) handleGetPosts(msg Message) {
	offset, _ := strconv.Atoi(msg.Message)
	posts := client.wsServer.getPosts(offset)
	message := Message{Posts: posts}
	if offset == 0 {
		message.Action = GetPostsAction
	} else {
		message.Action = MorePostsAction
	}
	client.send <- message.encode()
}

func (client *Client) handleGetCategoryPosts(msg Message) {
	offset, _ := strconv.Atoi(msg.Message)
	posts := client.wsServer.getCategoryPosts(msg.Target, offset)
	message := Message{Posts: posts}
	if offset == 0 {
		message.Action = GetCategoryPostsAction
	} else {
		message.Action = MorePostsAction
	}
	client.send <- message.encode()
}

func (client *Client) handleGetUserPosts(msg Message) {
	offset, _ := strconv.Atoi(msg.Message)
	posts := client.wsServer.getUserPosts(offset, client.GetName())
	message := Message{Posts: posts}
	if offset == 0 {
		message.Action = GetUserPostsAction
	} else {
		message.Action = MorePostsAction
	}
	client.send <- message.encode()
}

func (client *Client) handleGetPost(msg Message) {
	postId, _ := strconv.Atoi(msg.Message)
	post := client.wsServer.getPost(postId)
	message := Message{
		Posts:  []models.Post{post},
		Action: GetPostAction,
	}
	client.send <- message.encode()
}

func (client *Client) handleNewComment(msg Message) {
	postId, _ := strconv.Atoi(msg.Target)
	client.wsServer.addNewComment(postId, msg.Message, client.GetName())
	comment := client.wsServer.getLastComment(postId)
	message := Message{
		Comments: []models.Comment{comment},
		Action:   NewCommentAction,
	}
	client.send <- message.encode()
}

func (client *Client) handleGetComments(msg Message) {
	offset, _ := strconv.Atoi(msg.Message)
	postId, _ := strconv.Atoi(msg.Target)
	comments := client.wsServer.getComments(offset, postId)
	message := Message{Comments: comments}
	if offset == 0 {
		message.Action = GetCommentsAction
	} else {
		message.Action = MoreCommentsAction
	}

	client.send <- message.encode()
}

/* -------------------------------- // helper ------------------------------- */
/* -------- construct chat id based un usernames in ascii value order ------- */
func getChatId(sender, reciver string) string {
	chatId := ""
	if strings.Compare(sender, reciver) < 0 {
		chatId += sender + ":" + reciver
	} else {
		chatId += reciver + ":" + sender
	}
	return chatId
}

/* ------- /ServeWs handles websocket requests from clients requests. ------- */
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	// get the User from the context
	userCtxValue := r.Context().Value(auth.UserContextKey)
	if userCtxValue == nil {
		log.Println("Not authenticated")
		return
	}
	username := userCtxValue.(string)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Allow all connections
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer, username)

	go client.write()
	go client.read()

	wsServer.register <- client
}

/* --- main function for hadling reading msg -> communication with WS conn -- */
func (client *Client) read() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		client.handleNewMessage(jsonMessage)
	}

}

/* --- main function for hadling writing msg -> communication with WS conn -- */
func (client *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
