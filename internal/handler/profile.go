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

	err = db.DB.QueryRow(query, cookie.Value).Scan(&username, &userID)
	if err != nil {
		log.Println("Session not found or expired:", err)
		http.Redirect(w, req, "/users/login", http.StatusSeeOther)
		return
	}

	log.Println("Logged in user:", username)

	user, err := getUser(userID)
	log.Printf("%v", user)
	if err != nil {
		log.Printf("Error retrieving online users : %s", err.Error())
		ErrorPageHandler(w, "Error retrieving online users", http.StatusInternalServerError)
		return
	}

	data := db.PageData{
		HomePath:        homePagePath,
		Logo:            logPath,
		IsAuthenticated: isAuthenticated,
		Username:        username,
		HeaderCSS:       headerCSS,
		User:            user,
		ProfileCSS:      profilecss,
	}
	RenderTemplate(w, "profile", data)
}

func getUser(userID int) (db.User, error) {
	query := "SELECT UserID, NickName, Age, FirstName, LastName, Gender, Username, Email, RegistrationDate, is_online FROM User WHERE UserID = ?"
	row := db.DB.QueryRow(query, userID)

	var user db.User
	if err := row.Scan(&user.UserID, &user.NickName, &user.Age, &user.FirstName, &user.LastName, &user.Gender, &user.Username, &user.Email, &user.RegistrationDate, &user.IsOnline); err != nil {
		return db.User{}, err
	}

	return user, nil
}
