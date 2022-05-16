package auth

import (
	"context"
	"net/http"
	"real-time-forum/handlers"
)

type contextKey string

const UserContextKey = contextKey("user")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string `json:"id"`
}

func (user User) GetId() string {
	return user.Id
}
func (user User) GetName() string {
	return user.Username
}
func (user User) GetPassword() string {
	return user.Password
}

type Middleware struct {
	UserHandler *handlers.UserHandler
}

func (middle *Middleware) AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieFromWeb, err := r.Cookie("session-id")
		if err != nil {
			return
		}
		cookieValue := cookieFromWeb.Value
		if err != nil || len(cookieValue) == 0 {
			return
		}
		// Validate cookie
		foundUser := middle.UserHandler.FindUserBySessionId(cookieValue)
		if foundUser == "" {
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, foundUser)
		f(w, r.WithContext(ctx))
	})
}
