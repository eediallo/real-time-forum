package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

// ProfilePage renders the user's profile page.
func ProfilePage(w http.ResponseWriter, req *http.Request) {

	isAuthenticated := true

	cookie, err := req.Cookie("session_id")
	if err != nil {
		log.Println("No session_id cookie found:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	var username string
	query := `
		SELECT
				u.Username
			FROM 
				Session AS s
			INNER JOIN
				User AS u
			ON
				s.UserID = u.UserID
			WHERE 
				SessionID = ?`

	err = db.DB.QueryRow(query, cookie.Value).Scan(&username)
	if err != nil {
		log.Println("Session not found or expired:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	log.Println("Logged in user:", username)

	data := struct {
		HomePath        string
		Logo            string
		IsAuthenticated bool
		Username        string
		HeaderCSS       string
	}{HomePath: homePagePath, Logo: logPath, IsAuthenticated: isAuthenticated, Username: username, HeaderCSS: headerCSS}
	log.Printf("header css %s\n", data.HeaderCSS)

	RenderTemplate(w, "profile", data)
}
