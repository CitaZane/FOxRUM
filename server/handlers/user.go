package handlers

import (
	"database/sql"
	"log"
	"real-time-forum/models"
)

/* ---------------------------------- User ---------------------------------- */
type User struct {
	Id          string `json:"-"`
	Name        string `json:"name"`
	Password    string `json:"-"`
	Online      bool   `json:"online"`
	LastMsgTime string `json:"lastMsgTime"`
}

func (user *User) GetId() string {
	return user.Id
}

func (user *User) GetName() string {
	return user.Name
}
func (user *User) GetPassword() string {
	return user.Password
}
func (user *User) IsOnline() {
	user.Online = true
}
func (user *User) IsOffline() {
	user.Online = false
}
func (user *User) SetMessageTime(time string) {
	user.LastMsgTime = time
}

/* --------------------- user- handler / db interaction --------------------- */
type UserHandler struct {
	Db *sql.DB
}

/* ----------------------------- insert into db ----------------------------- */
func (handler *UserHandler) AddUser(user models.User, password, firstName, lastName, email, gender, birth string) {
	stmt, err := handler.Db.Prepare("INSERT INTO users(user_id, username, password, first_name, last_name, email, gender, birth_date) values(?,?,?,?,?,?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(user.GetId(), user.GetName(), password, firstName, lastName, email, gender, birth)
	checkErr(err)
}
func (handler *UserHandler) AddUserSID(username string, sessionId string) {
	stmt, err := handler.Db.Prepare("UPDATE users SET (session_id) = ? WHERE username = (?)")
	checkErr(err)

	_, err = stmt.Exec(sessionId, username)
	checkErr(err)
}

/* ----------------------------- remove / delete ---------------------------- */
func (handler *UserHandler) RemoveUser(user models.User) {
	stmt, err := handler.Db.Prepare("DELETE FROM users WHERE user_id = ?")
	checkErr(err)

	_, err = stmt.Exec(user.GetId())
	checkErr(err)
}

/* ----------------------------------- get ---------------------------------- */
func (handler *UserHandler) FindUserById(ID string) models.User {

	row := handler.Db.QueryRow("SELECT user_id, username FROM users where user_id = ? LIMIT 1", ID)

	var user User

	if err := row.Scan(&user.Id, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &user
}

func (handler *UserHandler) FindUserID(user models.User) string {

	row := handler.Db.QueryRow("SELECT user_id FROM users where username = ? LIMIT 1", user.GetName())
	var id string
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return id
}

func (handler *UserHandler) FindUserByUsername(username string) models.User {
	row := handler.Db.QueryRow("SELECT user_id, username, password FROM users where username = ? LIMIT 1", username)
	var user User
	if err := row.Scan(&user.Id, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}
	return &user
}

func (handler *UserHandler) FindUserByEmail(email string) *User {
	row := handler.Db.QueryRow("SELECT user_id, username, password FROM users where email = ? LIMIT 1", email)
	var user User
	if err := row.Scan(&user.Id, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}
	return &user
}
func (handler *UserHandler) FindUserBySessionId(sessionId string) string {
	row := handler.Db.QueryRow("SELECT  username FROM users where session_id = ? LIMIT 1", sessionId)
	var username string
	if err := row.Scan(&username); err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return username
}

func (handler *UserHandler) GetAllUsers() []models.User {
	rows, err := handler.Db.Query("SELECT user_id, username FROM users")

	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		user.Online = false
		users = append(users, &user)
	}

	return users
}

func (handler *UserHandler) GetUserJoinDate(username string) string {
	row := handler.Db.QueryRow("SELECT  created_at FROM users where username = ? LIMIT 1", username)
	var date string
	if err := row.Scan(&date); err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		panic(err)
	}
	return date
}
