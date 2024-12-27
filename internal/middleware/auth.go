package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

type contextKey string

const (
	userIDKey   contextKey = "userID"
	usernameKey contextKey = "username"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			log.Println("No session_id cookie found in middleware:", err)
			http.Redirect(w, r, loginPath, http.StatusSeeOther)
			return
		}

		var userID int
		var username sql.NullString
		err = db.DB.QueryRow(userByUsernameAndUserIDQuery,
			cookie.Value).Scan(&userID, &username)
		if err != nil {
			log.Println("Session not found or expired in middleware:", err)
			http.Redirect(w, r, loginPath, http.StatusSeeOther)
			return
		}

		if !username.Valid {
			log.Println("Username is NULL in session for user ID:", userID)
			http.Redirect(w, r, loginPath, http.StatusSeeOther)
			return
		}

		log.Println("Authenticated user:", username.String)
		// ensures that the UserID and Username are shared only within this context.
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(ctx, usernameKey, username.String)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
