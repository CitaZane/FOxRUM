package auth

import (
	"net/http"
	"net/url"
	"time"
)

func SendCookie(sessionID string, w http.ResponseWriter) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "session-id",
		Value:    url.QueryEscape(sessionID),
		Path:     "/",
		HttpOnly: false,
		Expires:  expiration,
		MaxAge:   86400,
	}
	http.SetCookie(w, &cookie)
	cs := w.Header().Get("Set-Cookie")
	cs += "; SameSite=lax"
	w.Header().Set("Set-Cookie", cs)
}
