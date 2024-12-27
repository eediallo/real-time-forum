package handler

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/eediallo/real_time_forum/internal/utils"
	"github.com/gofrs/uuid"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var errorMessage string
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		nickName := r.FormValue("nickname")
		password := r.FormValue("password")

		if email == "" && nickName == "" {
			errorMessage += "Email or NickName is required. "
		}
		if password == "" {
			errorMessage += "Password is required. "
		}

		var(
			storedHash string
			userID int
			username string
		)

		err := db.DB.QueryRow(userByEmailQuery, email, nickName).Scan(&storedHash, &userID, &nickName,  &username )
		if err != nil {
			if err == sql.ErrNoRows {
				ErrorPageHandler(w, "User not found", nil)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if !utils.CheckPasswordHash(password, storedHash) {
			ErrorPageHandler(w, "Invalid password for user:", username)
			return
		}

		// Invalidate any existing session for this user
		_, err = db.DB.Exec(deleteSessionByUserIDQuery, userID)
		if err != nil {
			log.Println("Error deleting existing session:", err)
			http.Redirect(w, r, loginPath, http.StatusSeeOther)
			return
		}

		// create a new session
		sessionID, _ := uuid.NewV4()
		_, err = db.DB.Exec(createNewSessionQuery, sessionID.String(), userID, time.Now())
		if err != nil {
			log.Println("Error inserting session:", err)
			http.Redirect(w, r, loginPath, http.StatusSeeOther)
			return
		}

		cookie := http.Cookie{
			Name:     "session_id",
			Value:    sessionID.String(),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600, // 1 hour
			Secure:   true,
		}
		http.SetCookie(w, &cookie)

		log.Println("Session ID set for user:", username, "with session ID:", sessionID.String())

		http.Redirect(w, r, dashboardPath, http.StatusSeeOther)
		return
	}

	data := PageData{
		CssLoginPath: cssLoginPath,
		Logo:         logPath,
		HomePath:     homePagePath,
		LogoCSS:      cssLogoPath,
		ErrorMessage: errorMessage,
		LoginJS:      loginJSPath,
	}

	RenderTemplate(w, "login", data)
}
