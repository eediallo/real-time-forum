package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

// DashboardPage handles the retrieval of posts and their comments and renders the dashboard page.
// If the user is logged in, it shows the username; otherwise, it displays the page without user-specific data.
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

	posts, err := fetchPosts()
	if err != nil {
		log.Printf("Error retrieving posts : %s", err.Error())
		ErrorPageHandler(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	users, err := fetchOnlineUsers()
	log.Printf("%v", users)
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
		OnlineUsers:          users,
		PrivateMessageJS:     privateMessageJS,
	}

	RenderTemplate(w, "dashboard", data)
}

// CreatePosts retrieves all posts from the database along with their associated comments.
func fetchPosts() ([]db.Post, error) {
	postRows, err := db.DB.Query(postsWithDetailsQuery)
	if err != nil {
		return nil, err
	}
	defer postRows.Close()

	var posts []db.Post
	for postRows.Next() {
		var post db.Post
		err := postRows.Scan(&post.PostID, &post.Title, &post.Content, &post.CommentCount, &post.CreatedAt, &post.Category, &post.Username)
		if err != nil {
			return nil, err
		}

		// Fetch like and dislike counts for the post
		post.LikeCount, post.DislikeCount, _ = getLikeDislikeCounts(post.PostID)

		comments, err := fetchComments(post.PostID)
		if err != nil {
			return nil, fmt.Errorf("error copying colums into row ---CreatePosts--%s", err.Error())
		}

		// Fetch comment count for the post
		err = db.DB.QueryRow("SELECT COUNT(*) FROM comments WHERE PostID = ?", post.PostID).Scan(&post.CommentCount)
		if err != nil {
			return nil, err
		}

		post.Comments = comments
		posts = append(posts, post)
	}

	return posts, nil
}

// CreateComments retrieves all comments for a given post ID from the database.
func fetchComments(postID int) ([]db.Comment, error) {
	commentRows, err := db.DB.Query(commentsFromDBQuery, postID)
	if err != nil {
		return nil, err
	}
	defer commentRows.Close()

	var comments []db.Comment
	for commentRows.Next() {
		var comment db.Comment
		err := commentRows.Scan(&comment.CommentID, &comment.PostID, &comment.Content, &comment.LikeCount, &comment.DislikeCount, &comment.CreatedAt, &comment.Username)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func fetchOnlineUsers() ([]db.User, error) {
	rows, err := db.DB.Query("SELECT UserID, Username, is_online, LastMessage FROM User WHERE is_online = 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []db.User
	for rows.Next() {
		var user db.User
		var lastMessage sql.NullString
		if err := rows.Scan(&user.UserID, &user.Username, &user.IsOnline, &lastMessage); err != nil {
			return nil, err
		}
		user.LastMessage = lastMessage.String
		users = append(users, user)
	}

	return users, nil
}
