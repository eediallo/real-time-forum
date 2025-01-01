package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/eediallo/real_time_forum/internal/db"
)

type FormData struct {
	Title    string
	Content  string
	Category string
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		log.Println("Error parsing form data:", err)
		return
	}

	// Retrieve the session ID from the cookie
	cookie, err := req.Cookie("session_id")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		log.Println("No session found:", err)
		return
	}

	// Retrieve the user ID and username from the session
	var userID int
	var username string
	err = db.DB.QueryRow(`
		SELECT 
			s.UserID,
			u.Username
		FROM 
			Session AS s
		INNER JOIN
			User AS u
		ON
			s.UserID = u.UserID
		WHERE 
			SessionID = ?`, cookie.Value).Scan(&userID, &username)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		log.Println("Invalid session:", err)
		return
	}

	// Get form values
	title := req.FormValue("title")
	content := req.FormValue("content")
	category := req.FormValue("category")

	// Insert the post into the database
	createdAT := time.Now()
	_, err = db.DB.Exec(`
	INSERT INTO 
		Post (UserID, Title, Content, CreatedAt, Category)
	VALUES 
		(?, ?, ?, ?, ?)`,
		userID, title, content, createdAT, category)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		log.Println("Error creating post:", err)
		return
	}
	http.Redirect(w, req, "/dashboard", http.StatusSeeOther)
}
