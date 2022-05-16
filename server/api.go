package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/auth"
	"real-time-forum/handlers"

	uuid "github.com/satori/go.uuid"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string

	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Gender          string `json:"gender"`
	Birth           string `json:"birth"`
	ConfirmPassword string `json:"confirm"`
	Email           string `json:"email"`
	Online          bool
	Time            string
}

func (user Credentials) GetId() string {
	return user.Id
}
func (user Credentials) GetName() string {
	return user.Username
}
func (user Credentials) GetPassword() string {
	return user.Password
}
func (user Credentials) IsOnline() {
	user.Online = true
}
func (user Credentials) IsOffline() {
	user.Online = false
}
func (user Credentials) SetMessageTime(time string) {
	user.Time = time
}

type API struct {
	UserHandler *handlers.UserHandler
}

func (api *API) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer func() { // recovers panic
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	w = configHeader(w)
	var user Credentials
	// Try to decode the JSON request to a LoginUser
	json.NewDecoder(r.Body).Decode(&user)
	if len(user.Username) == 0 {
		return
	} // catch empty requests

	/* ---------------------------- login validation ---------------------------- */
	username := ""
	emailErr := auth.ValidateEmail(user.Username)
	// loging in via email
	if emailErr.Message == "" {
		// find user in database
		dbUser := api.UserHandler.FindUserByEmail(user.Username)
		if dbUser == nil {
			respondWithError(w)
			return
		}
		err := auth.ComparePassword(user.Password, dbUser.Password)
		if err != nil {
			respondWithError(w)
			return
		}
		username = dbUser.Name
	} else { //logging in via username
		usernameErr := auth.ValidateUsername(user.Username)
		if usernameErr.Message != "" { //username error
			respondWithError(w)
			return
		}
		// find by username in database
		dbUser := api.UserHandler.FindUserByUsername(user.Username)
		if dbUser == nil {
			respondWithError(w)
			return
		}
		err := auth.ComparePassword(user.Password, dbUser.GetPassword())
		if err != nil {
			respondWithError(w)
			return
		}
		username = dbUser.GetName()
	}

	/* ---------------------------- login successful ---------------------------- */
	/* ---------------------------- configure cookie ---------------------------- */
	// Create sid
	sessionID := generateUUID()
	auth.SendCookie(sessionID, w)
	// Add sid to db
	api.UserHandler.AddUserSID(username, sessionID)
}

func (api *API) HandleSignup(w http.ResponseWriter, r *http.Request) {
	defer func() { // recovers panic
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	w = configHeader(w)
	var user Credentials
	json.NewDecoder(r.Body).Decode(&user) // Try to decode the JSON request to a user
	if len(user.Username) == 0 {
		return
	} // catch empty requests

	/* --------------- here should happen some back end validation -------------- */
	resp := []auth.ErrorResponse{}
	// validate all fields

	firstnameErr := auth.ValidateFirstname(user.FirstName)
	if firstnameErr.Message != "" {
		resp = append(resp, firstnameErr)
	}

	lastnameErr := auth.ValidateLastname(user.LastName)
	if lastnameErr.Message != "" {
		resp = append(resp, lastnameErr)
	}

	genderErr := auth.ValidateGender(user.Gender)
	if genderErr.Message != "" {
		resp = append(resp, genderErr)
	}

	birthErr := auth.ValidateBirth(user.Birth)
	if birthErr.Message != "" {
		resp = append(resp, birthErr)
	}

	emailErr := auth.ValidateEmail(user.Email)
	if emailErr.Message != "" {
		resp = append(resp, emailErr)
	} else {
		if api.UserHandler.FindUserByEmail(user.Email) != nil {
			emailErr = auth.ErrorResponse{Message: "Email already used", Property: "email"}
			resp = append(resp, emailErr)
		}
	}

	usernameErr := auth.ValidateUsername(user.Username)
	if usernameErr.Message != "" {
		resp = append(resp, usernameErr)
	} else {
		if api.UserHandler.FindUserByUsername(user.Username) != nil {
			usernameErr = auth.ErrorResponse{Message: "Username already taken", Property: "username"}
			resp = append(resp, usernameErr)
		}
	}

	passwordErr := auth.ValidatePassword(user.Password, user.ConfirmPassword)
	if passwordErr.Message != "" {
		resp = append(resp, passwordErr)
	}

	/* ------------------------- if some errors respond ------------------------- */
	if len(resp) != 0 {
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	/* ---------- validation successful continue with saving user in db --------- */
	hashedPwd, _ := auth.GeneratePassword(user.Password)

	user.Password = hashedPwd
	user.Id = generateUUID()
	api.UserHandler.AddUser(user, user.Password, user.FirstName, user.LastName, user.Email, user.Gender, user.Birth)
	sessionID := generateUUID()
	api.UserHandler.AddUserSID(user.Username, sessionID)
	// Create sid
	auth.SendCookie(sessionID, w)
}

func configHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	return w
}
func generateUUID() string {
	sessionToken := uuid.NewV4().String()
	return sessionToken
}

/* --------------- for logging in responds with standard error -------------- */
/* -------------------- to not revel too much information ------------------- */
func respondWithError(w http.ResponseWriter) {
	resp := []auth.ErrorResponse{}
	loginErr := auth.ErrorResponse{Message: "Username / e-mail or password wrong", Property: "username"}
	resp = append(resp, loginErr)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
