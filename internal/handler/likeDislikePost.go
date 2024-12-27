package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eediallo/real_time_forum/internal/db"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	handleLikeDislike(w, r, true)
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	handleLikeDislike(w, r, false)
}

func handleLikeDislike(w http.ResponseWriter, r *http.Request, isLike bool) {
	userID, err := getUserIDFromSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Check if the user has already liked/disliked the post
	var existingLikeDislikeID int
	var existingIsLike bool
	err = db.DB.QueryRow("SELECT LikeDislikeID, IsLike FROM LikeDislike WHERE UserID = ? AND PostID = ?", userID, postID).Scan(&existingLikeDislikeID, &existingIsLike)
	if err == sql.ErrNoRows {
		// Insert new like/dislike
		_, err := db.DB.Exec("INSERT INTO LikeDislike (UserID, PostID, IsLike) VALUES (?, ?, ?)", userID, postID, isLike)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	} else if err == nil {
		// Update existing like/dislike
		if existingIsLike != isLike {
			_, err := db.DB.Exec("UPDATE LikeDislike SET IsLike = ? WHERE LikeDislikeID = ?", isLike, existingLikeDislikeID)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
		}
	} else {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Respond with the updated like/dislike counts
	likeCount, dislikeCount, err := getLikeDislikeCounts(postID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	response := map[string]int{"likes": likeCount, "dislikes": dislikeCount}
	json.NewEncoder(w).Encode(response)
}

func getUserIDFromSession(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0, err
	}

	var userID int
	err = db.DB.QueryRow("SELECT UserID FROM Session WHERE SessionID = ?", cookie.Value).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func getLikeDislikeCounts(postID int) (int, int, error) {
	var likeCount, dislikeCount int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE PostID = ? AND IsLike = 1", postID).Scan(&likeCount)
	if err != nil {
		return 0, 0, err
	}
	err = db.DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE PostID = ? AND IsLike = 0", postID).Scan(&dislikeCount)
	if err != nil {
		return 0, 0, err
	}
	return likeCount, dislikeCount, nil
}
