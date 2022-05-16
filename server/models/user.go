package models

/* ---------------------- f-ns that interact with user ---------------------- */
type User interface {
	GetId() string
	GetName() string
	IsOnline()
	IsOffline()
	SetMessageTime(time string)
	GetPassword() string 
}

/* ----------------------- f-ns that interact with db ----------------------- */
type UserHandlers interface {
	AddUser(user User, password, firstName, lastName, email, gender, birth string)
	RemoveUser(user User)
	FindUserById(ID string) User
	GetAllUsers() []User
	FindUserID(user User) string
	GetUserJoinDate(username string) string
	FindUserByUsername(username string) User
}
