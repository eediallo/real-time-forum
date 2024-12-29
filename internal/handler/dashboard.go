package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

// DashboardPage handles the retrieval of posts and their comments and renders the dashboard page.
func DashboardPage(w http.ResponseWriter, r *http.Request) {
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
			ErrorPageHandler(w, "Session not found or expired:", nil)
			username = ""
			return
		}
	}

	posts, err := db.FetchPosts()
	if err != nil {
		log.Printf("Error retrieving posts : %s", err.Error())
		ErrorPageHandler(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	users, err := db.FetchAllUsers()
	if err != nil {
		log.Printf("Error retrieving online users : %s", err.Error())
		ErrorPageHandler(w, "Error retrieving online users", http.StatusInternalServerError)
		return
	}

	data := db.PageData{
		Username:             username,
		HeaderCSS:            headerCSS,
		Posts:                posts,
		DashboardCSS:         dashboardCSS,
		CommentJS:            commentJS,
		GoogleIcons:          googleIcons,
		Logo:                 logPath,
		HomePath:             homePagePath,
		LikeDislike:          likeDislike,
		IsAuthenticated:      isAuthenticated,
		LikeDislikeCommentJS: likeDislikeCommentJsPath,
		FilterJS:             filterJsPath,
		WS:                   wsPath,
		Users:                users,
		PrivateMessageJS:     privateMessageJS,
		MainJS:               mainjs,
	}

	RenderTemplate(w, "dashboard", data)
}
