package middleware

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

func CheckUserAuthentication(r *http.Request) (string, bool, error) {
	var username string
	var userID int
	isAuthenticated := false

	// Try to retrieve the session ID from the cookie
	cookie, err := r.Cookie("session_id")
	if err == nil {
		err = db.DB.QueryRow(sessionIDfromCookieQuery, cookie.Value).Scan(&userID, &username)
		isAuthenticated = true
		if err != nil {
			log.Printf("Session not found or expired in dashboard : %s", err.Error())
			return "", false, err
		}
	}

	return username, isAuthenticated, nil
}
