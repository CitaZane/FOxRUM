package main

import (
	"encoding/json"
	"log"
	"real-time-forum/models"
)

/* ----------- action variable used in go side / string for client ---------- */
/* ---------------------------------- chat ---------------------------------- */
const NewChatMessageAction = "new-chat-message"         // sending new message
const JoinChatAction = "join-chat"                      // opening chat
const ChatMessageRequestAction = "chat-message-request" //requesting more messages
const TypeAction = "type"
/* ------------------------------- user lists ------------------------------- */
const GetAllUsersAction = "get-all-users"       //geting all users from db
const GetOnlineUsersAction = "get-online-users" //getting all online users
/* ------------------------------ user actions ------------------------------ */
const GetUsernameAction = "get-username"
const UserLogoutAction = "user-logout" //user left
const UserJoinedAction = "user-join"   // respond to login or signup
const GetUserStatAction = "user-stats" // get statisctic for user
const ThrowOutAction = "throw-out"
/* ---------------------------------- post ---------------------------------- */
const GetCategoriesAction = "get-categories" // get list of categories
const NewPostAction = "new-post"
const GetPostsAction = "get-posts"
const GetCategoryPostsAction = "get-category-posts"
const MorePostsAction = "more-posts" // consquetive post request
const GetUserPostsAction = "get-user-posts"
const GetPostAction = "get-post" // get single post

/* -------------------------------- comments -------------------------------- */
const NewCommentAction = "new-comment"
const GetCommentsAction = "get-comments"
const MoreCommentsAction = "more-comments"

type Message struct {
	Action       string            `json:"action"`  //msg request action
	Message      string            `json:"message"` //The actual message
	Target       string            `jsob:"target"`  // a target for chat message / chat id / category
	Sender       models.User       `json:"sender"`
	Users        []models.User     `json:"users"` //for sending user lists
	Stats        map[string]string `json:"stats"` // for sending user statistics
	ChatMessages []models.Message  `json:"chatmessages"`
	Categories   []string          `json:"categories"` // for sending list of caegories
	Posts        []models.Post     `json:"posts"`
	Comments     []models.Comment  `json:"comments"`
}

// encode method that can be called to create a json []byte object
//  that is ready to be send back to client
func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return json
}

// UnmarshalJSON custom unmarshal to create a Client instance for Sender
func (message *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	msg := &struct {
		Sender Client `json:"sender"`
		*Alias
	}{
		Alias: (*Alias)(message),
	}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	message.Sender = &msg.Sender
	return nil
}
