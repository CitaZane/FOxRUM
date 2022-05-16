package auth

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ErrorResponse struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

func ValidateFirstname(name string) ErrorResponse {
	if fieldEmpty(name) {
		return ErrorResponse{Message: "First name is required", Property: "fistName"}
	}
	if len(name) > 60 {
		return ErrorResponse{Message: "Name too long", Property: "fistName"}
	}
	pattern := regexp.MustCompile("^\\w+$")
	if !pattern.MatchString(name) {
		return ErrorResponse{Message: "Invalid characters", Property: "fistName"}
	}
	return ErrorResponse{Message: ""}
}

func ValidateLastname(lastName string) ErrorResponse {
	if fieldEmpty(lastName) {
		return ErrorResponse{Message: "Last name is required", Property: "lastName"}
	}
	if len(lastName) > 60 {
		return ErrorResponse{Message: "Last name too long", Property: "lastName"}
	}
	pattern := regexp.MustCompile("^[A-Za-z-]+$")
	if !pattern.MatchString(lastName) {
		return ErrorResponse{Message: "Invalid characters", Property: "lastName"}
	}
	return ErrorResponse{Message: ""}
}

func ValidateGender(gender string) ErrorResponse {
	if fieldEmpty(gender) {
		return ErrorResponse{Message: "Gender is required", Property: "gender"}
	}
	list := []string{"female", "male", "other", "none"}
	listContains := false
	for i := 0; i < len(list); i++ {
		if gender == list[i] {
			listContains = true
			break
		}
		if i == len(list)-1 && !listContains {
			return ErrorResponse{Message: "Choose from provided options", Property: "gender"}
		}
	}
	return ErrorResponse{Message: ""}
}

func ValidateBirth(birth string) ErrorResponse {
	if fieldEmpty(birth) {
		return ErrorResponse{Message: "Date of birth is required", Property: "birth"}
	}
	pattern := regexp.MustCompile("^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$")
	if !pattern.MatchString(birth) {
		return ErrorResponse{Message: "Date of birth not valid", Property: "birth"}
	}
	t := time.Now()
	year, _ := strconv.Atoi(strings.Split(birth, "-")[0])
	if t.Year() < year {
		return ErrorResponse{Message: "Date of birth not valid", Property: "birth"}
	}
	if t.Year() < year-100 {
		return ErrorResponse{Message: "Date of birth not valid", Property: "birth"}
	}
	return ErrorResponse{Message: ""}
}

func ValidateEmail(email string) ErrorResponse {
	if fieldEmpty(email) {
		return ErrorResponse{Message: "Email is required", Property: "email"}
	}
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !pattern.MatchString(email) {
		return ErrorResponse{Message: "Email not valid", Property: "email"}
	}
	return ErrorResponse{Message: ""}
}

func ValidateUsername(username string) ErrorResponse {
	if fieldEmpty(username) {
		return ErrorResponse{Message: "Username is required", Property: "username"}
	}
	pattern := regexp.MustCompile("^[A-Za-z0-9_]+$")
	if len(username) < 4 {
		return ErrorResponse{Message: "Username too short", Property: "username"}
	}
	if len(username) > 15 {
		return ErrorResponse{Message: "Username too long", Property: "username"}
	}
	if !pattern.MatchString(username) {
		return ErrorResponse{Message: "Invalid characters", Property: "username"}
	}
	return ErrorResponse{Message: ""}
}

func ValidatePassword(password, confirm string) ErrorResponse {
	if fieldEmpty(password) || fieldEmpty(confirm) {
		return ErrorResponse{Message: "Password is required", Property: "password"}
	}
	if password != confirm {
		return ErrorResponse{Message: "Passwords do not match", Property: "password"}
	}
	if len(password) < 6 {
		return ErrorResponse{Message: "Password too short", Property: "password"}
	}
	return ErrorResponse{Message: ""}
}

/* ---------------------- helper chack for epmty input ---------------------- */
func fieldEmpty(value string) bool {
	return len(value) == 0
}
