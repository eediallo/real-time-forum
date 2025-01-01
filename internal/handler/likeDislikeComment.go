package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/eediallo/real_time_forum/internal/db"
)

func LikeDislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromSession(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	commentID, err := strconv.Atoi(r.URL.Query().Get("comment_id"))
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	isLike, err := strconv.ParseBool(r.URL.Query().Get("is_like"))
	if err != nil {
		http.Error(w, "Invalid like/dislike value", http.StatusBadRequest)
		return
	}

	var existingLikeDislikeID int
	var existingIsLike bool
	err = db.DB.QueryRow("SELECT LikeDislikeID, IsLike FROM LikeDislike WHERE UserID = ? AND CommentID = ?", userID, commentID).Scan(&existingLikeDislikeID, &existingIsLike)
	if err == sql.ErrNoRows {
		_, err := db.DB.Exec("INSERT INTO LikeDislike (UserID, CommentID, IsLike) VALUES (?, ?, ?)", userID, commentID, isLike)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	} else if err == nil {
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

	updateCommentLikeDislikeCounts(commentID)

	likeCount, dislikeCount := getLikeDislikeCountsForComment(commentID)
	response := map[string]int{"likes": likeCount, "dislikes": dislikeCount}
	json.NewEncoder(w).Encode(response)
}

func updateCommentLikeDislikeCounts(commentID int) {
	var likeCount, dislikeCount int
	db.DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE CommentID = ? AND IsLike = 1", commentID).Scan(&likeCount)
	db.DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE CommentID = ? AND IsLike = 0", commentID).Scan(&dislikeCount)

	_, err := db.DB.Exec("UPDATE Comments SET LikeCount = ?, DislikeCount = ? WHERE CommentID = ?", likeCount, dislikeCount, commentID)
	if err != nil {
		log.Printf("Error updating like/dislike counts for comment %d: %v", commentID, err)
	}
}

func getLikeDislikeCountsForComment(commentID int) (int, int) {
	var likeCount, dislikeCount int
	db.DB.QueryRow("SELECT LikeCount, DislikeCount FROM Comments WHERE CommentID = ?", commentID).Scan(&likeCount, &dislikeCount)
	return likeCount, dislikeCount
}
