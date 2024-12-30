package handler

import (
	"log"
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/gorilla/mux"
)

func ProfilePage(w http.ResponseWriter, req *http.Request) {
	// Extract the username from the URL
	vars := mux.Vars(req)
	profileUsername := vars["username"]

	// Verify if the user is authenticated
	cookie, err := req.Cookie("session_id")
	if err != nil {
		log.Println("No session_id cookie found:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	// Retrieve the logged-in user's information
	var loggedInUsername string
	var userID int
	query := `
        SELECT
            u.Username,
            u.UserID
        FROM 
            Session AS s
        INNER JOIN
            User AS u
        ON
            s.UserID = u.UserID
        WHERE 
            SessionID = ?`

	err = db.DB.QueryRow(query, cookie.Value).Scan(&loggedInUsername, &userID)
	if err != nil {
		log.Println("Session not found or expired:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	// Log the logged-in user's information
	log.Println("Logged in user:", loggedInUsername)

	// Fetch the profile for the requested username
	user, err := db.GetUserByUsername(profileUsername)
	if err != nil {
		log.Printf("Error retrieving user profile for username '%s': %s", profileUsername, err.Error())
		ErrorPageHandler(w, "User not found", http.StatusNotFound)
		return
	}

	// Render the profile page with the fetched user data
	data := db.PageData{
		HomePath:        homePagePath,
		Logo:            logPath,
		IsAuthenticated: true,
		Username:        loggedInUsername, // The logged-in user's username
		HeaderCSS:       headerCSS,
		User:            user,
		ProfileCSS:      profilecss,
	}
	RenderTemplate(w, "profile", data)
}
