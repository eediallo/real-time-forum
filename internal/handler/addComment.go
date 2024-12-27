package handler

import (
	"net/http"
	"time"

	"github.com/eediallo/real_time_forum/internal/db"
)

func AddComment(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the session ID from the cookie
	cookie, err := req.Cookie("session_id")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	// Retrieve the user ID and username from the session
	var userID int
	var username string
	err = db.DB.QueryRow(sessionIDfromCookieQuery, cookie.Value).Scan(&userID, &username)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	// Retrieve form values
	postID := req.FormValue("post_id")
	content := req.FormValue("content")

	// Insert comment into the database
	_, err = db.DB.Exec(insertCommentToDBQuery, postID, userID, content, time.Now())
	if err != nil {
		http.Error(w, "Error inserting comment", http.StatusInternalServerError)
		return
	}

	// Update comment count in the post
	_, err = db.DB.Exec("UPDATE Post SET CommentCount = CommentCount + 1 WHERE PostID = ?", postID)
	if err != nil {
		ErrorPageHandler(w, "Error updating comment count", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, dashboardPath, http.StatusSeeOther)
}
