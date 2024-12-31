package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println("No session cookie found:", err)
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
		return
	}

	log.Println("Session ID from cookie:", cookie.Value)

	var userID int
	err = db.DB.QueryRow(`
      		SELECT
                s.UserID
            FROM 
                Session AS s
            INNER JOIN
                User AS u
            ON
                s.UserID = u.UserID
            WHERE 
                SessionID = ?`, cookie.Value).Scan(&userID)
	if err != nil {
		log.Println("Error retrieving user ID from session:", err)
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	log.Println("User ID from session:", userID)

	_, err = db.DB.Exec("DELETE FROM Session WHERE SessionID = ?", cookie.Value)
	if err != nil {
		log.Println("Error deleting session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("UPDATE User SET is_online = 0 WHERE UserID = ?", userID)
	if err != nil {
		log.Println("Error updating user online status:", err)
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
		return
	}

	cookie = &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	log.Println("User logged out successfully")

	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}
