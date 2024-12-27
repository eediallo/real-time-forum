package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

// ProfilePage renders the user's profile page.
func ProfilePage(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session_id")
	if err != nil {
		log.Println("No session_id cookie found:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	var username string
	err = db.DB.QueryRow("SELECT Username FROM Session WHERE SessionID = ?", cookie.Value).Scan(&username)
	if err != nil {
		log.Println("Session not found or expired:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	log.Println("Logged in user:", username)

	data := struct {
		Username  string
		HeaderCSS string
	}{Username: username, HeaderCSS: headerCSS}

	RenderTemplate(w, "profile", data)
}
