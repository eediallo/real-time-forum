package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/eediallo/real_time_forum/internal/middleware"
)

// DashboardPage handles the retrieval of posts and their comments and renders the dashboard page.
func DashboardPage(w http.ResponseWriter, r *http.Request) {

	// Try to retrieve the session ID from the cookie
	username, isAuthenticated, errAuth := middleware.CheckUserAuthentication(r)
	if errAuth != nil {
		log.Printf("Error authenticating user : %s", errAuth.Error())
		return
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
